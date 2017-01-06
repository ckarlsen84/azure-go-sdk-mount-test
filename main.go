package main

import (
	"fmt"
	"log"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"time"
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"strings"
	"bytes"
	"io/ioutil"
	"io"
	"os"
)

func inspectRequest() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			log.Printf("Request: method=%s url=%s\n", r.Method, r.URL)
			return p.Prepare(r)
		})
	}
}

func inspectResponse() autorest.RespondDecorator {
	return func(r autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(resp *http.Response) error {
			var body, b bytes.Buffer

			resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, &body))
			if err := resp.Write(&b); err != nil {
				return fmt.Errorf("Failed to write response: %v", err)
			}

			resp.Body = ioutil.NopCloser(&body)
			log.Printf("Response: status=%s method=%s url=%s body=%s\n", resp.Status, resp.Request.Method, resp.Request.URL, body.String())

			return r.Respond(resp)
		})
	}
}

func newClient() (compute.VirtualMachinesClient, error) {
	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID"),
		"AZURE_RESOURCE_GROUP":  os.Getenv("AZURE_RESOURCE_GROUP"),
		"AZURE_VM_NAME":         os.Getenv("AZURE_VM_NAME"),
		"AZURE_DISK_NAME":       os.Getenv("AZURE_DISK_NAME"),
		"AZURE_DISK_URI":        os.Getenv("AZURE_DISK_URI"), }

	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)


	client := compute.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])

	//client.PollingDelay = time.Second * 1
	client.Authorizer = spt
	client.RequestInspector = inspectRequest()
	client.ResponseInspector = inspectResponse()

	if err != nil {
		log.Fatalf("Error: %v", err)
		return client, err
	}

	return client, nil

}

func testAttachDisk(client compute.VirtualMachinesClient) {
	start := time.Now()
	rgName := os.Getenv("AZURE_RESOURCE_GROUP")
	vmName := os.Getenv("AZURE_VM_NAME")
	diskName := os.Getenv("AZURE_DISK_NAME")
	diskURI := os.Getenv("AZURE_DISK_URI")
	var lun int32 = 0
	cachingMode := compute.CachingTypes("none")

	vm, err := client.Get(rgName, vmName, "")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	disks := *vm.StorageProfile.DataDisks
	disks = append(disks,
		compute.DataDisk{
			Name: &diskName,
			Vhd: &compute.VirtualHardDisk{
				URI: &diskURI,
			},
			Lun:          &lun,
			Caching:      cachingMode,
			CreateOption: "attach",
		})

	newVM := compute.VirtualMachine{
		Location: vm.Location,
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			StorageProfile: &compute.StorageProfile{
				DataDisks: &disks,
			},
		},
	}



	_, err = client.CreateOrUpdate(rgName, vmName, newVM, nil)

	if err != nil {
		log.Printf("azure attach failed, err: %v", err)
		detail := err.Error()
		if strings.Contains(detail, "Code=\"AcquireDiskLeaseFailed\"") {
			// if lease cannot be acquired, immediately detach the disk and return the original error
			log.Printf("failed to acquire disk lease, try detach")
			//az.DetachDiskByName(diskName, diskURI, nodeName)
		}
	} else {
		log.Printf("azure attach succeeded")
	}

	elapsed := time.Since(start)

	log.Printf("Disk attach took %s", elapsed)
}

func testDetachDisk(client compute.VirtualMachinesClient) (*compute.VirtualMachine, error) {
	start := time.Now()
	rgName := os.Getenv("AZURE_RESOURCE_GROUP")
	vmName := os.Getenv("AZURE_VM_NAME")
	diskName := os.Getenv("AZURE_DISK_NAME")
	diskURI := os.Getenv("AZURE_DISK_URI")

	vm, err := client.Get(rgName, vmName, "")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	disks := *vm.StorageProfile.DataDisks
	for i, disk := range disks {
		if (disk.Name != nil && diskName != "" && *disk.Name == diskName) || (disk.Vhd.URI != nil && diskURI != "" && *disk.Vhd.URI == diskURI) {
			// found the disk
			log.Printf("detach disk: name %q uri %q", diskName, diskURI)
			disks = append(disks[:i], disks[i+1:]...)
			break
		}
	}

	newVM := compute.VirtualMachine{
		Location: vm.Location,
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			StorageProfile: &compute.StorageProfile{
				DataDisks: &disks,
			},
		},
	}

	_, err = client.CreateOrUpdate(rgName, vmName, newVM, nil)

	if err != nil {
		log.Fatalf("azure detach failed, err: %v", err)
	} else {
		log.Printf("azure detach succeeded")
	}

	elapsed := time.Since(start)

	log.Printf("Disk deattach took %s", elapsed)

	return &vm, nil
}

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatalf("Could not create client: %v", err)
	}

	log.Print("Start testAttachDisk2")
	testAttachDisk(client)
	log.Print("Start testDetachDisk")
	testDetachDisk(client)
}
