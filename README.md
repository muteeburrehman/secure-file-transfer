# Secure File Transfer

# RSA Key Generation

This project requires RSA keys for encryption and decryption. Follow the steps below to generate and extract the required keys using OpenSSL.

## Generating RSA Keys

1. **Generate a 2048-bit Private Key**  
   Run the following command to create the private key:
   ```bash
   openssl genrsa -out private.pem 2048

2. **After generating the private key, extract public key** 
   Run the following command to extract the public key:
   ```bash 
   openssl rsa -in private.pem -outform PEM -pubout -out public.pem

3. **Install Dependencies**
    This project requires the following Go dependencies:
    ```bash
    gorilla/mux: A router and dispatcher for HTTP requests.
    jwt-go: A Go library for creating and verifying JWT tokens for authentication.
 