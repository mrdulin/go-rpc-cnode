const request = require('request');

const options = {
  url: "http://localhost:3000/rpc",
  method: "post",
  headers: {
    "content-type": "application/json"
  },
  json: true,
  body: {"jsonrpc": "1.0", "id": "1", "method": "UserService.GetUserByLoginname", "params": [{Loginname: "mrdulin"}] }
};

request(options, (error, response, body) => {
  if (error) {
    console.error('An error has occurred: ', error);
  } else {
    console.log('Post successful: response: ', body);
  }
});