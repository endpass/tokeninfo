# tokeninfo
Simple api to return info about Ethereum ERC20 tokens

This program serves an API based on a JSON file containing information about tokens and a directory
of images with their logos.

## Configuration
Most configuration is done through environment variables. These environment
variables are mandatory.

**TOKEN_LIST** Path to a JSON file containing information about tokens.
**TOKEN_IMAGE_DIR** Directory with images for token logos. File name is the
token address.

## API Endpoints
All endpoints return JSON responses.

### /api/v1/tokens
Returns an array of objects with information about all tokens

### /api/v1/token/{symbol}
Returns an object with information about a specific token
