package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"temporal-dsl/tdsl"
)

func main() {
	var dslWorkflow tdsl.Workflow

	dslWorkflow = createDSLWorkflow()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        "tdsl_" + uuid.New(),
		TaskQueue: "tdsl",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, tdsl.SimpleDSLWorkflow, dslWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

}

func createDSLWorkflow() tdsl.Workflow {

	activity := tdsl.ActivityInvocation{
		Name:      "SampleActivity1",
		Arguments: []string{"arg1", "arg2"},
		Result:    "result1",
	}

	statement := tdsl.Statement{
		Activity: &activity,
	}

	sequence := tdsl.Sequence{
		Elements: []*tdsl.Statement{
			&statement,
		},
	}

	root := tdsl.Statement{
		Sequence: &sequence,
	}

	workflow := tdsl.Workflow{
		Root: root,
	}

	return workflow

}
