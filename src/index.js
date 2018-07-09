const schedule = require('node-schedule');
const { exec } = require('child_process');
const fs = require('fs');
const path = require('path');

console.log('=== schedule commit start, commit every ===');

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

function renameSync() {
  const pathA = path.resolve(__dirname, './test.txt');
  const pathB = path.resolve(__dirname, './test-1.txt');
  if (fs.existsSync(pathA)) {
    fs.renameSync(pathA, pathB);
  } else {
    fs.renameSync(pathB, pathA);
  }
}

// const rule = '0 9 * * *';
const rule = '* 1 * * * *';
const j = schedule.scheduleJob(rule, () => {
  console.log('start to commit');
  console.log('next commit date: ', j.nextInvocation());
  renameSync();
  gitpush();
});

console.log(j.nextInvocation());
