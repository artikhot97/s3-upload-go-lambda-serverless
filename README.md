# s3-upload-go-lambda-serverless
This repo for upload image on s3 using lambda serverless 

Firstly you have to setup serverless 
1. Create dirctory and open in vs code
2. Install serverless 

Get all the templates
  sls create --help
For this function
  sls create --template aws-go-dep --path upload-s3
  
Write Code for operation 

Upload lambda function
make deploy
![Screenshot from 2022-09-21 16-45-48](https://user-images.githubusercontent.com/61792772/191743367-1f6442af-4470-4c8d-966b-8f0b44f19a68.png)

On sucessfull deploy

![Screenshot from 2022-09-21 17-28-28](https://user-images.githubusercontent.com/61792772/191743810-2455b4f0-761c-48ca-b317-f6a866522f3d.png)

 
Encoder URLs
https://base64.guru/converter/encode/image

https://base64.guru/converter/decode/image

Example :

![Screenshot from 2022-09-21 20-44-38](https://user-images.githubusercontent.com/61792772/191742957-0da0e479-4b9b-4a7d-8bbe-b40fb919375a.png)

