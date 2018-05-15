const schedule = require('node-schedule');
const { exec } = require('child_process');

// const rule = '3 * * * * *';
// const j = schedule.scheduleJob(rule, () => {
//   console.log('test');
// });

const command = "git add . && git commit -m 'commit' && git push origin master";
function gitpush(callback) {
  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
  });
}

gitpush();
