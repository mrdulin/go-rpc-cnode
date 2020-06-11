function promiseAll<T>(ps: Array<Promise<T>>): Promise<T[]> {
  return new Promise((resolve, reject) => {
    let count = 0;
    const values: T[] = [];
    for (let i = 0; i < ps.length; i++) {
      Promise.resolve(ps[i])
        .then((value) => {
          count--;
          values[i] = value;
          if (count === 0) {
            resolve(values);
          }
        })
        .catch(reject);
      count++;
    }
  });
}

export { promiseAll };
