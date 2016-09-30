function rq(i) {
  return new Promise((resolve, reject) =>
    setTimeout(() => {
      console.log('11');
      if (i === 2) {
        return reject('error happened');
      }
      console.log('------');
      resolve(22);
    }, 1000)
  );
}

export { rq };
