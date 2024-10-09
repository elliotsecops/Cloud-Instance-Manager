package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var client *ec2.Client
var dryRun bool

func init() {
	dryRun = os.Getenv("DRY_RUN") == "true"
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("Error al cargar la configuración de AWS: %v", err)
	}
	client = ec2.NewFromConfig(cfg)
}

func crearInstancia(amiID string) (string, error) {
	result, err := client.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		ImageId:      aws.String(amiID),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
		DryRun:       &dryRun,
	})
	if err != nil {
		return "", err
	}
	if dryRun {
		return "i-00000000000000000", nil // ID de instancia ficticio para pruebas
	}
	return *result.Instances[0].InstanceId, nil
}

func listarInstancias() ([]types.Instance, error) {
	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		DryRun: &dryRun,
	})
	if err != nil {
		return nil, err
	}
	if dryRun {
		// Devolver instancias ficticias para pruebas
		return []types.Instance{
			{InstanceId: aws.String("i-00000000000000001"), State: &types.InstanceState{Name: types.InstanceStateNameRunning}},
			{InstanceId: aws.String("i-00000000000000002"), State: &types.InstanceState{Name: types.InstanceStateNameStopped}},
		}, nil
	}
	var instances []types.Instance
	for _, reservation := range result.Reservations {
		instances = append(instances, reservation.Instances...)
	}
	return instances, nil
}

func iniciarInstancia(instanceID string) error {
	_, err := client.StartInstances(context.TODO(), &ec2.StartInstancesInput{
		InstanceIds: []string{instanceID},
		DryRun:      &dryRun,
	})
	return err
}

func detenerInstancia(instanceID string) error {
	_, err := client.StopInstances(context.TODO(), &ec2.StopInstancesInput{
		InstanceIds: []string{instanceID},
		DryRun:      &dryRun,
	})
	return err
}

func terminarInstancia(instanceID string) error {
	_, err := client.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceID},
		DryRun:      &dryRun,
	})
	return err
}

func main() {
	fmt.Printf("Ejecutando en modo de prueba (DryRun): %v\n", dryRun)

	instanceID, err := crearInstancia("ami-12345678")
	if err != nil {
		log.Fatalf("Error al crear instancia: %v", err)
	}
	fmt.Printf("Instancia creada: %s\n", instanceID)

	instances, err := listarInstancias()
	if err != nil {
		log.Fatalf("Error al listar instancias: %v", err)
	}
	for _, instance := range instances {
		fmt.Printf("ID: %s, Estado: %s\n", *instance.InstanceId, instance.State.Name)
	}

	if err := iniciarInstancia(instanceID); err != nil {
		log.Fatalf("Error al iniciar instancia: %v", err)
	}
	fmt.Println("Instancia iniciada")

	if err := detenerInstancia(instanceID); err != nil {
		log.Fatalf("Error al detener instancia: %v", err)
	}
	fmt.Println("Instancia detenida")

	if err := terminarInstancia(instanceID); err != nil {
		log.Fatalf("Error al terminar instancia: %v", err)
	}
	fmt.Println("Instancia terminada")
}
