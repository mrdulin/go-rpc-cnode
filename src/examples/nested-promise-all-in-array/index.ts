import { asyncTask } from '../../lib/asyncTask';

const tasks: Array<Promise<any>> = [];

const taskQueue: Promise<any> = Promise.all([asyncTask('task A'), asyncTask('task B')]);
const taskC = asyncTask('task C');

tasks.concat(taskQueue);
tasks.push(taskC);

const label = 'task run';
console.time(label);
Promise.all([taskC, taskQueue])
  .then(results => {
    console.log('results: ', results);
    console.timeEnd(label);
  })
  .catch(console.error);
