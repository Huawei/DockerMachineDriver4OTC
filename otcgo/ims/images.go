package ims

import (
	"encoding/json"

	"github.com/docker/machine/libmachine/log"
)

type ImageList struct {
	File              string   `json:"file"`
	Owner             string   `json:"owner"`
	Id                string   `json:"id"`
	Size              int64    `json:"size"`
	Self              string   `json:"self"`
	Schema            string   `json:"schema"`
	Status            string   `sjon:"status"`
	Tags              []string `json:"tags"`
	Visibility        string   `json:"visibility"`
	Name              string   `json:"name"`
	Checksum          string   `json:"checksum"`
	Deleted           bool     `json:"deleted"`
	Protected         bool     `sjon:"protected"`
	Container_Format  string   `json:"container_format"`
	Min_Ram           uint32   `json:"min_ram"`
	Update_At         string   `json:"update_at"`
	Os_Bit            string   `json:"__os_bit"`
	Os_Version        string   `json:"__os_version"`
	Description       string   `json:"__description"`
	Disk_Format       string   `json:"disk_format"`
	IsRegistered      string   `json:"__isregistered"`
	Platform          string   `json:"__platform"`
	Os_Type           string   `json:"__os_type"`
	Min_Disk          uint32   `json:"min_disk"`
	Virtual_Env_Type  string   `json:"virtual_env_type"`
	Image_Source_Type string   `json:"__images_source_type"`
	ImageType         string   `json:"__imagetype"`
	Create_At         string   `json:"create_at"`
	Virtual_Size      uint32   `json:"virtual_size"`
	Deleted_At        string   `json:"deleted_at"`
	OriginalImageName string   `json:"__originalimagename"`
	Backup_Id         string   `json:"__backup_id"`
	ProductCode       string   `json:"__productcode"`
	Image_Location    string   `json:"__image_location"`
	Image_Size        string   `json:"__image_size"`
	Data_Origin       string   `json:"__data_origin"`
}

type ImageResponse struct {
	Images []ImageList `json:"images"`
}

func (client *Client) ListImages(region string) (imgList *ImageResponse, err error) {
	var imageList ImageResponse

	//compose uri string
	uri := "/cloudimages"

	respbytes, err := client.Do(region, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	//unmarshal response body data into image list
	if err := json.Unmarshal(respbytes, &imageList); err != nil {
		return nil, err
	}

	log.Debugf("image list are: %v", imageList)
	return &imageList, nil
}
