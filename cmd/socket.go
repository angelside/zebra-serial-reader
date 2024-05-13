package main

import (
	"bufio"
	"net"
	"time"
    "strings"
	//"fmt"
    //"os"
)

func getSocketAnswer(ip string, command string) (string, error) {
    // Create a dialer with the timeout
    dialer := net.Dialer{Timeout: 3 * time.Second}

    // Make a connection
    conn, err := dialer.Dial("tcp", ip+":9100") // FIXME: hardcoded printer port
	if err != nil {
		return "", err
	}
    /*
    if err != nil {
        println("Dial failed: ", err.Error())
        os.Exit(1)
    }
    */
    defer conn.Close()

    rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

    //println("Sending string request")
    //_, err = rw.WriteString("! U1 getvar \"device.unique_id\"\r\n")
    _, err = rw.WriteString(command)
    if err != nil {
		return "", err
	}
    /*
    if err != nil {
        println("Error getting status.", err)
    }
    */
    //println("Characters sent: ", n) // replaced with _

    err = rw.Flush()
    if err != nil {
        //println("Error on flush", err)
        return "", err
    }

    err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
    if err != nil {
        //println("Set Deadline failed: ", err)
        return "", err
    }

    response, err := rw.ReadString('\n')
    if err != nil {
        //println("Read error: ", err.Error()) // Read error:  read tcp 172.18.210.76:51429->172.18.69.228:9100: i/o timeout
    }
    //println("Response: ", response)

    //return response, nil
    return strings.Trim(response, `"`), nil
}
