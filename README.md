# Matrices
Matrices is a simple RESTful API Application that accepts HTTP JSON POST Request of two matrices, performs multiplication of those matrices and returns the resulting matrix as a JSON Response

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/atharva-shinde/matrices)
![GitHub Repo stars](https://img.shields.io/github/stars/atharva-shinde/matrices?style=flat)
![GitHub License](https://img.shields.io/github/license/atharva-shinde/matrices?style=flat)
![GitHub Issues](https://img.shields.io/github/issues/atharva-shinde/matrices)
![GitHub Pull Requests](https://img.shields.io/github/issues-pr/atharva-shinde/matrices)


---

## Table of Contents
- [Prerequisites](#prerequisites)
- [Local Installation](#local-installation)
    - [Clone  the Repository](#clone-the-repository)
    - [Run the HTTP server](#run-the-http-server)
    - [Sending a HTTP API Request](#sending-n-http-api-request)
- [Makefile Targets](#makefile-targets)
- [LICENSE](#license)



## Prerequisites:
1. We will be publishing the application image to [quay.io](https://quay.io) image registry. Export your quay.io username as an environment variable `QUAY`
    
    `export QUAY="<your_quay_username>"`

    example: `export QUAY="atharva"`

2. Export Openshift Container Platform Login and Password Credentials as an environment variable `OCP_LOGIN`

    `export OCP_LOGIN="<ocp credentials>"`
    
    example: `export OCP_LOGIN="oc login https://api.cluster.com:6443 -u kubeadmin -p password`
"

---

## Local Installation

### Clone the Repository
`git clone git@github.com:Atharva-Shinde/matrices.git`

### Run the HTTP server

1. `cd matrices`
2. `make run`

The server is now listening on port `8080` 

### Sending a HTTP API Request
Open a new terminal and send a HTTP request

```
curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d '{"matrices":[{"rows":3,"columns":3,"data":[[1, 1, 1],[2, 2, 2],[3, 3, 3]]},{"rows": 3,"columns": 3,"data": [[1, 1, 1],[2, 2, 2],[3, 3, 3]]}]}'
```

JSON Request:
```json
{
				"matrices": [
					{
						"rows": 3,
						"columns": 3,
						"data": [
							[1,1,1],
							[2,2,2],
							[3,3,3]
						]
					},
					{
						"rows": 3,
						"columns": 3,
						"data": [
							[1,1,1],
							[2,2,2],
							[3,3,3]
						]
					}
				]
			}
```

API response:
```json
{
  "rows": 3,
  "columns": 3,
  "data": [
    [6,6,6],
    [12,12,12],
    [18,18,18]
  ]
}
```

---

## Makefile targets

1. Help

    `make help`

2. Build a binary of the application

    `make binary`

3. Create an image of the application

    `make image`

4. Run the application by executing the binary

    `make run`

5. Publish the application image to quay.io

    `make publish`

6. Deploy the application to an OCP cluster

    `make deploy-docker` deploys the application to OCP using docker build strategy

    `make deploy-source` deploys the application to OCP using S2I(Source to Image) strategy

7. Test the application

    `make test` 

---

## LICENSE

matrices is licensed under the [MIT LICENSE](https://github.com/Atharva-Shinde/matrices/blob/main/LICENSE)