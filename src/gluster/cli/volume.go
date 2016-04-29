package cli

import (
	"encoding/xml"

	"fmt"
)

// VolumeOption object
type VolumeOption struct {
	Name  string `xml:"name" json:"name,omitempty"`
	Value string `xml:"value" json:"value,omitempty"`
}

// Transport type
type Transport string

// Volume object
type Volume struct {
	Name            string         `xml:"name" json:"name"`
	ID              string         `xml:"id" json:"id"`
	Status          string         `xml:"statusStr" json:"status"`
	Type            string         `xml:"typeStr" json:"type"`
	Bricks          []Brick        `xml:"bricks>brick" json:"bricks"`
	NumBricks       int            `xml:"brickCount" json:"num_bricks"`
	DistCount       int            `xml:"distCount" json:"dist_count"`
	ReplicaCount    int            `xml:"replicaCount" json:"replica_count"`
	StripeCount     int            `xml:"stripeCount" json:"stripe_count"`
	ArbiterCount    int            `xml:"arbiterCount" json:"arbiter_count"`
	DisperseCount   int            `xml:"disperseCount" json:"disperse_count"`
	RedundancyCount int            `xml:"redundancyCount" json:"redundancy_count"`
	TransportRaw    Transport      `xml:"transport" json:"transport"`
	Options         []VolumeOption `xml:"options>option" json:"options"`
}

// Ports object of Bricks
type Ports struct {
	TCP  string `xml:"tcp" json:"tcp,omitempty"`
	Rdma string `xml:"rdma" json:"rdma,omitempty"`
}

// Brick object
type Brick struct {
	Name       string `xml:"name" json:"name,omitempty"`
	UUID       string `xml:"hostUuid" json:"host_id,omitempty"`
	HostUUID   string `xml:"peerid" json:"-"`
	Hostname   string `xml:"hostname" json:"hostname,omitempty"`
	Path       string `xml:"path" json:"path,omitempty"`
	StatusRaw  int    `xml:"status" json:"-"`
	Online     bool   `json:"online,omitempty"`
	Ports      *Ports `xml:"ports" json:"ports,omitempty"`
	Pid        string `xml:"pid" json:"pid,omitempty"`
	SizeTotal  string `xml:"sizeTotal" json:"size_total,omitempty"`
	SizeFree   string `xml:"sizeFree" json:"size_free,omitempty"`
	Device     string `xml:"device" json:"device,omitempty"`
	BlockSize  string `xml:"blockSize" json:"block_size,omitempty"`
	MntOptions string `xml:"mntOptions" json:"mnt_options,omitempty"`
	FsName     string `xml:"fsName" json:"fs_name,omitempty"`
}

// BricksStatus from Gluster status output
type BricksStatus struct {
	XMLName xml.Name `xml:"cliOutput"`
	List    []Brick  `xml:"volStatus>volumes>volume>node"`
}

// Volumes - List of Volume objects
type Volumes struct {
	XMLName xml.Name `xml:"cliOutput"`
	List    []Volume `xml:"volInfo>volumes>volume"`
}

// VolListVolumes - List of Volumes from vol list command
type VolListVolumes struct {
	XMLName xml.Name `xml:"cliOutput"`
	List    []string `xml:"volList>volume"`
}

// CreateOptions - Options to Create Volume
type CreateOptions struct {
	Bricks            []string `json:"bricks"`
	ReplicaCount      int      `json:"replica"`
	StripeCount       int      `json:"stripe"`
	ArbiterCount      int      `json:"arbiter"`
	DisperseCount     int      `json:"disperse"`
	DisperseDataCount int      `json:"disperse-data"`
	RedundancyCount   int      `json:"disperse-redundancy"`
	Transport         string   `json:"transport"`
	AllowRootDir      bool     `json:"allow_root_dir"`
	ReuseBricks       bool     `json:"reuse_bricks"`
}

// VolumeCreate is a func to create Gluster Volume
func VolumeCreate(volname string, bricks []string, options CreateOptions) error {
	// volume create <NEW-VOLNAME> [stripe <COUNT>] [replica <COUNT> [arbiter <COUNT>]]
	// [disperse [<COUNT>]] [disperse-data <COUNT>] [redundancy <COUNT>]
	// [transport <tcp|rdma|tcp,rdma>] <NEW-BRICK>?<vg_name>... [force]
	// - create a new volume of specified type with mentioned bricks
	// TODO Validate the Inputs and Options
	cmd := []string{"volume", "create", volname}
	if options.ReplicaCount != 0 {
		cmd = append(cmd, "replica", fmt.Sprintf("%d", options.ReplicaCount))
	}
	if options.StripeCount != 0 {
		cmd = append(cmd, "stripe", fmt.Sprintf("%d", options.StripeCount))
	}
	if options.ArbiterCount != 0 {
		cmd = append(cmd, "arbiter", fmt.Sprintf("%d", options.ArbiterCount))
	}
	if options.DisperseCount != 0 {
		cmd = append(cmd, "disperse", fmt.Sprintf("%d", options.DisperseCount))
	}
	if options.DisperseDataCount != 0 {
		cmd = append(cmd, "disperse-data", fmt.Sprintf("%d", options.DisperseDataCount))
	}
	if options.RedundancyCount != 0 {
		cmd = append(cmd, "redundancy", fmt.Sprintf("%d", options.RedundancyCount))
	}
	if options.Transport != "" {
		cmd = append(cmd, "transport", options.Transport)
	}

	cmd = append(cmd, bricks...)

	if options.AllowRootDir || options.ReuseBricks {
		// Limitation in Gluster 1.0, Can't differenciate between
		// multiple Flags. Common option "force"
		cmd = append(cmd, "force")
	}

	return ExecuteCmd(cmd)
}

// VolumeStart is a func to start a Gluster Volume
func VolumeStart(volname string, force bool) error {
	cmd := []string{"volume", "start", volname}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

// VolumeStop is a func to stop a Gluster Volume
func VolumeStop(volname string, force bool) error {
	cmd := []string{"volume", "stop", volname}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

// VolumeDelete is a func to delete a Gluster Volume
func VolumeDelete(volname string) error {
	cmd := []string{"volume", "delete", volname}
	return ExecuteCmd(cmd)
}

// VolumeOptSet is a func to set option of a Gluster Volume
func VolumeOptSet(volname string, key string, value string) error {
	// TODO: Handle Multiple options
	cmd := []string{"volume", "set", volname, key, value}
	return ExecuteCmd(cmd)
}

// VolumeOptReset is a func to reset option of a Gluster Volume
func VolumeOptReset(volname string, key string, force bool) error {
	cmd := []string{"volume", "reset", volname}
	if key != "" {
		cmd = append(cmd, key)
	}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

// VolumeOptGet is a func to get Gluster Volume options
func VolumeOptGet(volname string, key string) ([]VolumeOption, error) {
	if key == "" {
		// If key is empty then run Volume info and return the options
		var q Volumes
		cmd := []string{"volume", "info", volname}
		data, err := ExecuteCmdXML(cmd)
		if err != nil {
			return []VolumeOption{}, err
		}
		xmlerr := xml.Unmarshal(data, &q)
		if xmlerr != nil {
			return []VolumeOption{}, xmlerr
		}
		return q.List[0].Options, nil
	} else if key == "all" {
		// If key is "all" run volume get <VOL> all and return output
		// TODO: Open issue with xml output
		return []VolumeOption{}, nil
	} else {
		// If key is set then get for that key
		// TODO: Open issue with xml output
		return []VolumeOption{}, nil
	}
}

// VolumeLogRotate is a utility func to initiate log rotate on a Gluster Volume
func VolumeLogRotate(volname string, brick string) error {
	// TODO: brick is mandate, wrong vol help
	cmd := []string{"volume", "log", volname, "rotate"}
	if brick != "" {
		cmd = append(cmd, brick)
	}
	return ExecuteCmd(cmd)
}

// VolumeRestart is a utility func to restart Gluster Volume
func VolumeRestart(volname string, force bool) error {
	errStop := VolumeStop(volname, force)
	if errStop != nil {
		return errStop
	}
	return VolumeStart(volname, force)
}

// VolumeInfo is a utility func to get Gluster Volume information
// by running gluster volume info command
func VolumeInfo(volname string) ([]Volume, error) {
	var q Volumes
	cmd := []string{"volume", "info"}
	if volname != "" {
		cmd = append(cmd, volname)
	}
	data, err := ExecuteCmdXML(cmd)
	if err != nil {
		return []Volume{}, err
	}
	xmlerr := xml.Unmarshal(data, &q)
	if xmlerr != nil {
		return []Volume{}, xmlerr
	}
	return q.List, nil
}

// VolumeStatus is a utility func to get Volume status by running gluster
// volume status and info command. This func merges the output from both
// Volume info command and Volume status command to show offline node status
func VolumeStatus(volname string) ([]Volume, error) {
	var tmpBrickStatus = make(map[string]Brick)

	var bricks BricksStatus
	vol := "all"
	if volname != "" {
		vol = volname
	}
	cmd := []string{"volume", "status", vol, "detail"}
	data, err := ExecuteCmdXML(cmd)
	if err != nil {
		return []Volume{}, err
	}
	xmlerr := xml.Unmarshal(data, &bricks)
	if xmlerr != nil {
		return []Volume{}, xmlerr
	}

	// Create hashmap to lookup and merge later
	for _, b := range bricks.List {
		name := b.Hostname + ":" + b.Path
		b.Name = name
		if b.StatusRaw == 0 {
			b.Online = false
		} else {
			b.Online = true
		}
		b.UUID = b.HostUUID
		tmpBrickStatus[name] = b
	}

	var volumes Volumes
	cmd1 := []string{"volume", "info"}
	if volname != "" {
		cmd1 = append(cmd1, volname)
	}
	data1, err1 := ExecuteCmdXML(cmd1)
	if err1 != nil {
		return []Volume{}, err
	}
	xmlerr1 := xml.Unmarshal(data1, &volumes)
	if xmlerr1 != nil {
		return []Volume{}, xmlerr
	}

	for idx, v := range volumes.List {
		for idx1, b := range v.Bricks {
			if brickData, ok := tmpBrickStatus[b.Name]; ok {
				volumes.List[idx].Bricks[idx1] = brickData
			} else {
				// Brick node is not online, Set default values
				volumes.List[idx].Bricks[idx1].Online = false
				volumes.List[idx].Bricks[idx1].Ports = &Ports{TCP: "N/A", Rdma: "N/A"}
				volumes.List[idx].Bricks[idx1].Pid = "N/A"
				volumes.List[idx].Bricks[idx1].SizeTotal = "N/A"
				volumes.List[idx].Bricks[idx1].SizeFree = "N/A"
				volumes.List[idx].Bricks[idx1].Device = "N/A"
				volumes.List[idx].Bricks[idx1].BlockSize = "N/A"
				volumes.List[idx].Bricks[idx1].MntOptions = "N/A"
				volumes.List[idx].Bricks[idx1].FsName = "N/A"
			}
		}
	}

	return volumes.List, nil
}

// VolumeList by running gluster volume list command
func VolumeList() ([]string, error) {
	var q VolListVolumes
	cmd := []string{"volume", "list"}
	data, err := ExecuteCmdXML(cmd)
	if err != nil {
		return []string{}, err
	}
	xmlerr := xml.Unmarshal(data, &q)
	if xmlerr != nil {
		return []string{}, xmlerr
	}
	return q.List, nil
}

// VolumeBarrierEnable to enable IO barrier
func VolumeBarrierEnable(volname string) error {
	cmd := []string{"volume", "barrier", volname, "enable"}
	return ExecuteCmd(cmd)
}

// VolumeBarrierDisable to disable IO barrier
func VolumeBarrierDisable(volname string) error {
	cmd := []string{"volume", "barrier", volname, "disable"}
	return ExecuteCmd(cmd)
}
