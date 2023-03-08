package eventtrigger


type EventTriggerEventProcessors struct {
	// aws_eventbridge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/event_trigger#aws_eventbridge EventTrigger#aws_eventbridge}
	AwsEventbridge *EventTriggerEventProcessorsAwsEventbridge `field:"optional" json:"awsEventbridge" yaml:"awsEventbridge"`
}

