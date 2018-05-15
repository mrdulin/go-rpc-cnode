const schedule = require('node-schedule');
const { exec } = require('child_process');

// const rule = '3 * * * * *';
// const j = schedule.scheduleJob(rule, () => {
//   console.log('test');
// });

function gitpush(callback) {
  exec('npm run git:push', (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
    process.exit(1);
  });
}

gitpush();
