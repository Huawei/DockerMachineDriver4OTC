/*
 *Copyright 2015 Huawei Technologies Co., Ltd. All rights reserved.
 *	   eSDK is licensed under the Apache License, Version 2.0 (the "License");
 *	   you may not use this file except in compliance with the License.
 *	   You may obtain a copy of the License at
 *
 *	       http://www.apache.org/licenses/LICENSE-2.0
 *
 *
 *	   Unless required by applicable law or agreed to in writing, software
 *	   distributed under the License is distributed on an "AS IS" BASIS,
 *	   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *	   See the License for the specific language governing permissions and
 *	   limitations under the License.
 */
package imsModules

import (
	"fmt"
)

// URI parameters
type ListCloudImagesReqEx struct {
	absolutePath map[string]string
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) Init()
 * @brief		Initialize
 * @param[in]
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) Init() {
	listCloudImagesReqEx.absolutePath = make(map[string]string)
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetImageType(__imagetype string)
 * @brief		Specifies the image type.
 * @param[in]	__imagetype: Specifies the image type. The value can be gold,
							 which indicates a public image, or private,
							which indicates a private image.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetImageType(__imagetype string) {
	listCloudImagesReqEx.absolutePath["__imagetype"] = __imagetype
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetProtected(protected bool)
 * @brief		Specifies whether the image is protected.
 * @param[in]	protected: Specifies whether the image is protected.
						   The value must be set to True when you query public images.
						   This parameter is optional when you query private images.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetProtected(protected bool) {
	listCloudImagesReqEx.absolutePath["protected"] = fmt.Sprintf("%t", protected)
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetVisibility(visibility string)
 * @brief		Specifies whether the image can be seen by other tenants.
 * @param[in]	visibility: Specifies whether the image can be seen by other tenants.
							The value can be private or public.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetVisibility(visibility string) {
	listCloudImagesReqEx.absolutePath["visibility"] = visibility
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetOwner(owner string)
 * @brief		Specifies the tenant to which the image belongs.
 * @param[in]	owner: Specifies the tenant to which the image belongs.
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetOwner(owner string) {
	listCloudImagesReqEx.absolutePath["owner"] = owner
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetId(id string)
 * @brief		Specifies the image ID.
 * @param[in]	id: Specifies the image ID.
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetId(id string) {
	listCloudImagesReqEx.absolutePath["id"] = id
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetStatus(status string)
 * @brief		Specifies the image status.
 * @param[in]	status: Specifies the image status. The value can be active, queued,
					saving, deleted, or killed. An image can be used only
					when it is in the active state.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetStatus(status string) {
	listCloudImagesReqEx.absolutePath["status"] = status
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetName(name string)
 * @brief		Specifies the image name.
 * @param[in]	name: Specifies the image name. The image name is a string of 128
					  characters consisting of letters, digits, underscores (_), and spaces.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetName(name string) {
	listCloudImagesReqEx.absolutePath["name"] = name
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetContainerFormat(container_format string)
 * @brief		Specifies the container type.
 * @param[in]	container_format: Specifies the container type.
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetContainerFormat(container_format string) {
	listCloudImagesReqEx.absolutePath["container_format"] = container_format
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetDiskFormat(disk_format string)
 * @brief		Specifies the image format.
 * @param[in]	disk_format: Specifies the image format. Only the zvhd format is supported.
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetDiskFormat(disk_format string) {
	listCloudImagesReqEx.absolutePath["disk_format"] = disk_format
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetMinRam(min_ram string)
 * @brief		Specifies the minimum memory size (MB) required for running the image.
 * @param[in]	min_ram: Specifies the minimum memory size (MB) required for running the image.
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetMinRam(min_ram string) {
	listCloudImagesReqEx.absolutePath["min_ram"] = min_ram
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetMinDisk(min_disk string)
 * @brief		Specifies the minimum disk space (GB) required for running the image.
 * @param[in]	min_disk: Specifies the minimum disk space (GB) required for running the image.
 * @param[out]
 * @return
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetMinDisk(min_disk string) {
	listCloudImagesReqEx.absolutePath["min_disk"] = min_disk
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetMarker(marker string)
 * @brief		Specifies the start number from which images are queried.
 * @param[in]	marker: Specifies the start number from which images are queried.
						The value is the image ID.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetMarker(marker string) {
	listCloudImagesReqEx.absolutePath["marker"] = marker
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetLimit(limit int)
 * @brief		Specifies the number of images to be queried.
 * @param[in]	limit: Specifies the number of images to be queried.
					   The value is an integer.
					   All requests records can be queried by default.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetLimit(limit int) {
	listCloudImagesReqEx.absolutePath["limit"] = fmt.Sprintf("%d", limit)
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetOsBit(__os_bit string)
 * @brief		Specifies the system property of the operating system.
 * @param[in]	__os_bit: Specifies the system property of the operating system.
						  The value can be 32 or 64.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetOsBit(__os_bit string) {
	listCloudImagesReqEx.absolutePath["__os_bit"] = __os_bit
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetPlatform(__platform string)
 * @brief		Specifies the image platform type.
 * @param[in]	__platform: Specifies the image platform type.
							The value can be Windows, Ubuntu, RedHat, SUSE, CentOS, or Other.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetPlatform(__platform string) {
	listCloudImagesReqEx.absolutePath["__platform"] = __platform
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetSortKey(sort_key string)
 * @brief		Specifies the field for sorting the query results.
 * @param[in]	sort_key: Specifies the field for sorting the query results.
						  The value can be attributes of the image: name, container_format,
						  disk_format, status, id, and size. created_at is set for sorting
						  the query results by default.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetSortKey(sort_key string) {
	listCloudImagesReqEx.absolutePath["sort_key"] = sort_key
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetSortDir(sort_dir string)
 * @brief		Specifies whether the query results are sorted in ascending or descending order.
 * @param[in]	sort_dir: Specifies whether the query results are sorted in ascending or
						  descending order. The value can be asc or desc.
						  This parameter used together with parameter sort_key which indicates
						  that the query results are sorted in descending order by default.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetSortDir(sort_dir string) {
	listCloudImagesReqEx.absolutePath["sort_dir"] = sort_dir
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) SetOsType(__os_type string)
 * @brief		Specifies the operating system type.
 * @param[in]	__os_type: Specifies the operating system type.
						   The value can be Linux, Windows, or Other.
 * @param[out]
 * @return
*/
func (listCloudImagesReqEx *ListCloudImagesReqEx) SetOsType(__os_type string) {
	listCloudImagesReqEx.absolutePath["__os_type"] = __os_type
}

/*
 * @fn			func (listCloudImagesReqEx *ListCloudImagesReqEx) GetAbsolutePath() string
 * @brief		Return URI
 * @param[in]
 * @param[out]
 * @return		string
 */
func (listCloudImagesReqEx *ListCloudImagesReqEx) GetAbsolutePath() string {
	var strAbsolute string
	if len(listCloudImagesReqEx.absolutePath) > 0 {
		strAbsolute = "?"
		for key, _ := range listCloudImagesReqEx.absolutePath {
			strAbsolute += key + "=" + listCloudImagesReqEx.absolutePath[key] + "&"
		}
		strAbsolute = strAbsolute[:(len([]rune(strAbsolute)) - 1)]
	} else {
		strAbsolute = ""
	}

	return strAbsolute
}
