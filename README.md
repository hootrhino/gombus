# Go MeterBus

Golang Meter Bus Protocol Simple Implementation. The functionality is not comprehensive, but all common protocols are supported. The project is from: https://github.com/jonaz/gombus. The original project was not well-designed and contained bugs. This version is an improved edition after reorganizing the files.

## Example
### TCP
```go

	conn, err := gombus.OpenTCP("192.168.13.42:10001")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	frame := gombus.SndNKE(uint8(primaryID))
	fmt.Printf("sending nke: % x\n", frame)
	_, err = conn.Write(frame)
	if err != nil {
		logrus.Error(err)
		return
	}
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		logrus.Error(err)
		return
	}

	// frame := gombus.SetPrimaryUsingPrimary(0, 3)
	respFrame := &gombus.DecodedFrame{}
	lastFCB := true
	frames := 0
	for respFrame.HasMoreRecords() || frames == 0 {
		frames++
		// frame := gombus.SetPrimaryUsingPrimary(0, 3)
		frame = gombus.RequestUD2(uint8(primaryID))
		if !lastFCB {
			frame.SetFCB()
			frame.SetChecksum()
		}
		lastFCB = frame.C().FCB()

		fmt.Printf("sending: % x\n", frame)
		fmt.Printf("sending C: % b\n", frame.C())
		_, err = conn.Write(frame)
		if err != nil {
			logrus.Error(err)
			return
		}

		resp, err := gombus.ReadLongFrame(conn)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Printf("read: % x\n", resp)
		fmt.Printf("read C: % b\n", resp.C())

		respFrame, err = resp.Decode()
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Info("read total values: ", len(respFrame.DataRecords))
	}

	logrus.Info("read total frames: ", frames)
```
### Serial
```go

	conn, err := gombus.OpenSerial(serial.Config{
		Address:  "COM1",
		BaudRate: 9600,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
		Timeout:  3 * time.Second,
	})
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	frame := gombus.SndNKE(uint8(primaryID))
	fmt.Printf("sending nke: % x\n", frame)
	_, err = conn.Write(frame)
	if err != nil {
		logrus.Error(err)
		return
	}
	_, err = gombus.ReadSingleCharFrame(conn)
	if err != nil {
		logrus.Error(err)
		return
	}

	// frame := gombus.SetPrimaryUsingPrimary(0, 3)
	respFrame := &gombus.DecodedFrame{}
	lastFCB := true
	frames := 0
	for respFrame.HasMoreRecords() || frames == 0 {
		frames++
		// frame := gombus.SetPrimaryUsingPrimary(0, 3)
		frame = gombus.RequestUD2(uint8(primaryID))
		if !lastFCB {
			frame.SetFCB()
			frame.SetChecksum()
		}
		lastFCB = frame.C().FCB()

		fmt.Printf("sending: % x\n", frame)
		fmt.Printf("sending C: % b\n", frame.C())
		_, err = conn.Write(frame)
		if err != nil {
			logrus.Error(err)
			return
		}

		resp, err := gombus.ReadLongFrame(conn)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Printf("read: % x\n", resp)
		fmt.Printf("read C: % b\n", resp.C())

		respFrame, err = resp.Decode()
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Info("read total values: ", len(respFrame.DataRecords))
	}

	logrus.Info("read total frames: ", frames)
```