package server

import (
	"context"

	"githib.com/g41797/memphisgrpc/pb"
	"github.com/memphisdev/memphis.go"
)

type manager struct {
	bc *brokerConnector
	mc *memphis.Conn
}

func newManager(bc *brokerConnector) *manager {
	manager := new(manager)
	manager.bc = bc
	return manager
}

func (srv *manager) CreateStation(ctx context.Context, req *pb.CreateStationRequest) (*pb.Status, error) {

	mc, err := srv.bc.connect()
	if err != nil {
		text := err.Error()
		status := pb.Status{}
		status.Text = &text
		return &status, err
	}

	srv.mc = mc

	_, status := connCreateStation(srv.mc, req)

	return status, nil
}
func (srv *manager) DestroyStation(ctx context.Context, req *pb.DestroyStationRequest) (*pb.Status, error) {

	mc, err := srv.bc.connect()
	if err != nil {
		text := err.Error()
		status := pb.Status{}
		status.Text = &text
		return &status, err
	}

	srv.mc = mc

	status := connDestroyStation(srv.mc, req)

	return status, nil
}

func (srv *manager) clean() {
	if srv == nil {
		return
	}

	if srv.mc != nil {
		srv.mc.Close()
	}
}

func connCreateStation(conn *memphis.Conn, req *pb.CreateStationRequest) (*memphis.Station, *pb.Status) {

	status := pb.Status{}
	text := ""

	station := req.GetStation()
	if (station == nil) || (len(station.GetName()) == 0) {
		text = "empty station"
		status.Text = &text
		return nil, &status
	}

	sname := station.GetName()

	opts, err := stationOpts(req)

	if err != nil {
		text = err.Error()
		status.Text = &text
		return nil, &status
	}

	mst, err := conn.CreateStation(sname, opts...)
	if err != nil {
		text = err.Error()
		status.Text = &text
		return nil, &status
	}

	return mst, &status
}

func stationOpts(req *pb.CreateStationRequest) ([]memphis.StationOpt, error) {
	opts := make([]memphis.StationOpt, 0)
	return opts, nil
}

func connDestroyStation(conn *memphis.Conn, req *pb.DestroyStationRequest) *pb.Status {

	status := pb.Status{}
	text := ""

	station := req.GetStation()
	if (station == nil) || (len(station.GetName()) == 0) {
		text = "empty station"
		status.Text = &text
		return &status
	}

	sname := station.GetName()

	st, err := conn.CreateStation(sname)
	if err != nil {
		return &status
	}

	err = st.Destroy()
	if err != nil {
		text = err.Error()
		status.Text = &text
		return &status
	}

	return &status
}

/*
func connCreateProducer(conn *memphis.Conn, req *pb.CreateProducerRequest) (string, *pb.Status) {
	return "", nil
}

func connDestroyProducer(conn *memphis.Conn, req *pb.DestroyProducerRequest) *pb.Status {
	return nil
}

func connCreateConsumer(conn *memphis.Conn, req *pb.CreateConsumerRequest) (string, *pb.Status) {
	return "", nil
}

func connDestroyConsumer(conn *memphis.Conn, req *pb.DestroyConsumerRequest) *pb.Status {
	return nil
}
*/
