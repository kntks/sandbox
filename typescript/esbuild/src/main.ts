import dayjs from "dayjs";

function hoge(): string {
  return dayjs(new Date()).format("YYYY-MM-DD")
}

console.log(hoge())
export {}