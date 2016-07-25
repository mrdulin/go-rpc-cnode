function createUser() {
  return Promise.reject('create user failed.');
}

function publish() {
  return Promise.reject('publish message failed.');
}
async function main() {
  return createUser()
    .catch(error => {
      console.error(error);
      return publish();
    })
    .catch(error => {
      console.error(error);
    });
}

main();
