import { student_list_orch } from "../loadtest_api/student_list_orch.js";
import Options from "../option.config.js";

export const options = Options.const_arrive_time();

export default function () {
  student_list_orch();
}
