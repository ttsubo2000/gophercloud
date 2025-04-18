package images

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	th "github.com/ttsubo2000/gophercloud/testhelper"
	fakeclient "github.com/ttsubo2000/gophercloud/testhelper/client"
)

type imageEntry struct {
	ID   string
	JSON string
}

// HandleImageListSuccessfully test setup
func HandleImageListSuccessfully(t *testing.T) {

	images := make([]imageEntry, 3)

	images[0] = imageEntry{"cirros-0.3.4-x86_64-uec",
		`{
            "status": "active",
            "name": "cirros-0.3.4-x86_64-uec",
            "tags": [],
            "kernel_id": "e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4",
            "container_format": "ami",
            "created_at": "2015-07-15T11:43:35Z",
            "ramdisk_id": "8c64f48a-45a3-4eaa-adff-a8106b6c005b",
            "disk_format": "ami",
            "updated_at": "2015-07-15T11:43:35Z",
            "visibility": "public",
            "self": "/v2/images/07aa21a9-fa1a-430e-9a33-185be5982431",
            "min_disk": 0,
            "protected": false,
            "id": "07aa21a9-fa1a-430e-9a33-185be5982431",
            "size": 25165824,
            "file": "/v2/images/07aa21a9-fa1a-430e-9a33-185be5982431/file",
            "checksum": "eb9139e4942121f22bbc2afc0400b2a4",
            "owner": "cba624273b8344e59dd1fd18685183b0",
            "virtual_size": null,
            "min_ram": 0,
            "schema": "/v2/schemas/image"
        }`}
	images[1] = imageEntry{"cirros-0.3.4-x86_64-uec-ramdisk",
		`{
            "status": "active",
            "name": "cirros-0.3.4-x86_64-uec-ramdisk",
            "tags": [],
            "container_format": "ari",
            "created_at": "2015-07-15T11:43:32Z",
            "size": 3740163,
            "disk_format": "ari",
            "updated_at": "2015-07-15T11:43:32Z",
            "visibility": "public",
            "self": "/v2/images/8c64f48a-45a3-4eaa-adff-a8106b6c005b",
            "min_disk": 0,
            "protected": false,
            "id": "8c64f48a-45a3-4eaa-adff-a8106b6c005b",
            "file": "/v2/images/8c64f48a-45a3-4eaa-adff-a8106b6c005b/file",
            "checksum": "be575a2b939972276ef675752936977f",
            "owner": "cba624273b8344e59dd1fd18685183b0",
            "virtual_size": null,
            "min_ram": 0,
            "schema": "/v2/schemas/image"
        }`}
	images[2] = imageEntry{"cirros-0.3.4-x86_64-uec-kernel",
		`{
            "status": "active",
            "name": "cirros-0.3.4-x86_64-uec-kernel",
            "tags": [],
            "container_format": "aki",
            "created_at": "2015-07-15T11:43:29Z",
            "size": 4979632,
            "disk_format": "aki",
            "updated_at": "2015-07-15T11:43:30Z",
            "visibility": "public",
            "self": "/v2/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4",
            "min_disk": 0,
            "protected": false,
            "id": "e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4",
            "file": "/v2/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4/file",
            "checksum": "8a40c862b5735975d82605c1dd395796",
            "owner": "cba624273b8344e59dd1fd18685183b0",
            "virtual_size": null,
            "min_ram": 0,
            "schema": "/v2/schemas/image"
        }`}

	th.Mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		limit := 10
		var err error
		if r.FormValue("limit") != "" {
			limit, err = strconv.Atoi(r.FormValue("limit"))
			if err != nil {
				t.Errorf("Error value for 'limit' parameter %v (error: %v)", r.FormValue("limit"), err)
			}

		}

		marker := ""
		newMarker := ""

		if r.Form["marker"] != nil {
			marker = r.Form["marker"][0]
		}

		t.Logf("limit = %v   marker = %v", limit, marker)

		selected := 0
		addNext := false
		var imageJSON []string

		fmt.Fprintf(w, `{"images": [`)

		for _, i := range images {
			if marker == "" || addNext {
				t.Logf("Adding image %v to page", i.ID)
				imageJSON = append(imageJSON, i.JSON)
				newMarker = i.ID
				selected++
			} else {
				if strings.Contains(i.JSON, marker) {
					addNext = true
				}
			}

			if selected == limit {
				break
			}
		}
		t.Logf("Writing out %v image(s)", len(imageJSON))
		fmt.Fprintf(w, strings.Join(imageJSON, ","))

		fmt.Fprintf(w, `],
			    "next": "/images?marker=%s&limit=%v",
			    "schema": "/schemas/images",
			    "first": "/images?limit=%v"}`, newMarker, limit, limit)

	})
}

// HandleImageCreationSuccessfully test setup
func HandleImageCreationSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, `{
			"id": "e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
			"name": "Ubuntu 12.10",
			"protected": false,
			"tags": [
				"ubuntu",
				"quantal"
			]
		}`)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "queued",
			"name": "Ubuntu 12.10",
			"protected": false,
			"tags": ["ubuntu","quantal"],
			"container_format": "bare",
			"created_at": "2014-11-11T20:47:55Z",
			"disk_format": "qcow2",
			"updated_at": "2014-11-11T20:47:55Z",
			"visibility": "private",
			"self": "/v2/images/e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
			"min_disk": 0,
			"protected": false,
			"id": "e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
			"file": "/v2/images/e7db3b45-8db7-47ad-8109-3fb55c2c24fd/file",
			"owner": "b4eedccc6fb74fa8a7ad6b08382b852b",
			"min_ram": 0,
			"schema": "/v2/schemas/image",
			"size": 0,
			"checksum": "",
			"virtual_size": 0
		}`)
	})
}

// HandleImageCreationSuccessfullyNulls test setup
// JSON null values could be also returned according to behaviour https://bugs.launchpad.net/glance/+bug/1481512
func HandleImageCreationSuccessfullyNulls(t *testing.T) {
	th.Mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, `{
			"id": "e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
			"name": "Ubuntu 12.10",
			"protected": false,
			"tags": [
				"ubuntu",
				"quantal"
			]
		}`)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "queued",
			"name": "Ubuntu 12.10",
			"protected": false,
			"tags": ["ubuntu","quantal"],
			"container_format": "bare",
			"created_at": "2014-11-11T20:47:55Z",
			"disk_format": "qcow2",
			"updated_at": "2014-11-11T20:47:55Z",
			"visibility": "private",
			"self": "/v2/images/e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
			"min_disk": 0,
			"protected": false,
			"id": "e7db3b45-8db7-47ad-8109-3fb55c2c24fd",
			"file": "/v2/images/e7db3b45-8db7-47ad-8109-3fb55c2c24fd/file",
			"owner": "b4eedccc6fb74fa8a7ad6b08382b852b",
			"min_ram": 0,
			"schema": "/v2/schemas/image",
			"size": null,
			"checksum": null,
			"virtual_size": null
		}`)
	})
}

// HandleImageGetSuccessfully test setup
func HandleImageGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1bea47ed-f6a9-463b-b423-14b9cca9ad27", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "active",
			"name": "cirros-0.3.2-x86_64-disk",
			"tags": [],
			"container_format": "bare",
			"created_at": "2014-05-05T17:15:10Z",
			"disk_format": "qcow2",
			"updated_at": "2014-05-05T17:15:11Z",
			"visibility": "public",
			"self": "/v2/images/1bea47ed-f6a9-463b-b423-14b9cca9ad27",
			"min_disk": 0,
			"protected": false,
			"id": "1bea47ed-f6a9-463b-b423-14b9cca9ad27",
			"file": "/v2/images/1bea47ed-f6a9-463b-b423-14b9cca9ad27/file",
			"checksum": "64d7c1cd2b6f60c92c14662941cb7913",
			"owner": "5ef70662f8b34079a6eddb8da9d75fe8",
			"size": 13167616,
			"min_ram": 0,
			"schema": "/v2/schemas/image",
			"virtual_size": "None"
		}`)
	})
}

// HandleImageDeleteSuccessfully test setup
func HandleImageDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/1bea47ed-f6a9-463b-b423-14b9cca9ad27", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleImageUpdateSuccessfully setup
func HandleImageUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `[
			{
				"op": "replace",
				"path": "/name",
				"value": "Fedora 17"
			},
			{
				"op": "replace",
				"path": "/tags",
				"value": [
					"fedora",
					"beefy"
				]
			}
		]`)

		th.AssertEquals(t, "application/openstack-images-v2.1-json-patch", r.Header.Get("Content-Type"))

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": "da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
			"name": "Fedora 17",
			"status": "active",
			"visibility": "public",
			"size": 2254249,
			"checksum": "2cec138d7dae2aa59038ef8c9aec2390",
			"tags": [
				"fedora",
				"beefy"
			],
			"created_at": "2012-08-10T19:23:50Z",
			"updated_at": "2012-08-12T11:11:33Z",
			"self": "/v2/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
			"file": "/v2/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/file",
			"schema": "/v2/schemas/image",
			"owner": "",
			"min_ram": 0,
			"min_disk": 0,
			"disk_format": "",
			"virtual_size": 0,
			"container_format": ""
		}`)
	})
}

// HandlePutImageDataSuccessfully setup
func HandlePutImageDataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/file", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Unable to read request body: %v", err)
		}

		th.AssertByteArrayEquals(t, []byte{5, 3, 7, 24}, b)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleGetImageDataSuccessfully setup
func HandleGetImageDataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/file", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte{34, 87, 0, 23, 23, 23, 56, 255, 254, 0})
		th.AssertNoErr(t, err)
	})
}

// HandleCreateImageMemberSuccessfully setup
func HandleCreateImageMemberSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `{"member": "8989447062e04a818baf9e073fd04fa7"}`)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		    "created_at": "2013-09-20T19:22:19Z",
		    "image_id": "da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
		    "member_id": "8989447062e04a818baf9e073fd04fa7",
		    "schema": "/v2/schemas/member",
		    "status": "pending",
		    "updated_at": "2013-09-20T19:25:31Z"
			}`)

	})
}

// HandleCreateImageMemberInvalidVisibility setup for case when visibility=public
func HandleCreateImageMemberInvalidVisibility(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-invalid-visibility/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `{"member": "8989447062e04a818baf9e073fd04fa7"}`)
		w.WriteHeader(http.StatusForbidden)
	})

}

// HandleCreateImageMemberConflict setup for case when member is already image member
func HandleCreateImageMemberConflict(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-memberConflict/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `{"member": "8989447062e04a818baf9e073fd04fa7"}`)

		w.WriteHeader(http.StatusConflict)
	})
}

// HandleImageMemberList happy path setup
func HandleImageMemberList(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		    "members": [
		        {
		            "created_at": "2013-10-07T17:58:03Z",
		            "image_id": "da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
		            "member_id": "123456789",
		            "schema": "/v2/schemas/member",
		            "status": "pending",
		            "updated_at": "2013-10-07T17:58:03Z"
		        },
		        {
		            "created_at": "2013-10-07T17:58:55Z",
		            "image_id": "da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
		            "member_id": "987654321",
		            "schema": "/v2/schemas/member",
		            "status": "accepted",
		            "updated_at": "2013-10-08T12:08:55Z"
		        }
		    ],
		    "schema": "/v2/schemas/members"
		}`)
	})
}

// HandleImageMemberEmptyList happy path setup
func HandleImageMemberEmptyList(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		    "members": [],
		    "schema": "/v2/schemas/members"
		}`)
	})
}

// HandleImageMemberDetails setup
func HandleImageMemberDetails(t *testing.T) {
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members/8989447062e04a818baf9e073fd04fa7", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		    "status": "pending",
		    "created_at": "2013-11-26T07:21:21Z",
		    "updated_at": "2013-11-26T07:21:21Z",
		    "image_id": "da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
		    "member_id": "8989447062e04a818baf9e073fd04fa7",
		    "schema": "/v2/schemas/member"
		}`)
	})
}

// HandleImageMemberDeleteSuccessfully setup
func HandleImageMemberDeleteSuccessfully(t *testing.T) *CallsCounter {
	var counter CallsCounter
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members/8989447062e04a818baf9e073fd04fa7", func(w http.ResponseWriter, r *http.Request) {
		counter.Counter = counter.Counter + 1

		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
	return &counter
}

// HandleImageMemberDeleteByNonOwner setup
func HandleImageMemberDeleteByNonOwner(t *testing.T) *CallsCounter {
	var counter CallsCounter
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members/8989447062e04a818baf9e073fd04fa7", func(w http.ResponseWriter, r *http.Request) {
		counter.Counter = counter.Counter + 1

		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusForbidden)
	})
	return &counter
}

// HandleImageMemberUpdate setup
func HandleImageMemberUpdate(t *testing.T) *CallsCounter {
	var counter CallsCounter
	th.Mux.HandleFunc("/images/da3b75d9-3f4a-40e7-8a2c-bfab23927dea/members/8989447062e04a818baf9e073fd04fa7", func(w http.ResponseWriter, r *http.Request) {
		counter.Counter = counter.Counter + 1

		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `{"status": "accepted"}`)

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `{
		    "status": "accepted",
		    "created_at": "2013-11-26T07:21:21Z",
		    "updated_at": "2013-11-26T07:21:21Z",
		    "image_id": "da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
		    "member_id": "8989447062e04a818baf9e073fd04fa7",
		    "schema": "/v2/schemas/member"
		}`)
	})
	return &counter
}

// CallsCounter for checking if request handler was called at all
type CallsCounter struct {
	Counter int
}
