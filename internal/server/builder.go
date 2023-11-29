package server

import "githib.com/g41797/memphisgrpc/pb"

func createStationRequest(sname string) (req *pb.CreateStationRequest) {

	station := &pb.Station{Name: sname}
	abret := &pb.RetentionOpt_Abret{Abret: &pb.AckBasedRet{AckBased: false}}
	retopt := &pb.RetentionOpt{Retentions: abret}
	stopt := &pb.StorageOpt{StorageType: pb.StorageOpt_Disk}
	partopt := &pb.PartitionOpt{Number: 1}
	sopts := &pb.StationOpions{Part: partopt, Storage: stopt, Retention: retopt}

	return &pb.CreateStationRequest{Station: station, Options: sopts}
}

func destroyStationRequest(sname string) (req *pb.DestroyStationRequest) {

	station := &pb.Station{Name: sname}

	destst := &pb.DestroyStationRequest{Station: station}

	return destst
}

func startProduceRequest(sname, pname string) *pb.ProduceMessages {

	station := &pb.Station{Name: sname}
	producer := &pb.Producer{Name: pname}
	cpreq := &pb.CreateProducerRequest{Station: station, Producer: producer}
	start := &pb.ProduceMessages_Start{Start: cpreq}

	request := &pb.ProduceMessages{Data: start}

	return request
}

func produceRequest(hdrs map[string]string, body []byte) *pb.ProduceMessages {
	h := &pb.Headers{Headers: hdrs}
	msg := &pb.ProduceMessages_Msg{Msg: &pb.Msg{Headers: h, Body: body}}

	request := &pb.ProduceMessages{Data: msg}

	return request
}

func stopProduceRequest() *pb.ProduceMessages {
	stop := &pb.ProduceMessages_Stop{Stop: &pb.Stop{}}

	request := &pb.ProduceMessages{Data: stop}

	return request
}

func startConsumeRequest(sname, name string) *pb.ConsumeMessages {

	station := &pb.Station{Name: sname}
	consumer := &pb.Consumer{Name: name}
	req := &pb.CreateConsumerRequest{Station: station, Consumer: consumer}
	start := &pb.ConsumeMessages_Start{Start: req}

	request := &pb.ConsumeMessages{Data: start}

	return request
}

func stopConsumeRequest() *pb.ConsumeMessages {
	stop := &pb.ConsumeMessages_Stop{Stop: &pb.Stop{}}

	request := &pb.ConsumeMessages{Data: stop}

	return request
}
