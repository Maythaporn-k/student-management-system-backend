import { student_list_core } from "../loadtest_api/core/student_list_core.js";
import Options from "../../option.config.js";

export const thresholds_429 = {
  "http_req_failed{status:429}": ["rate<0.05"],
};

export const options = Options.const_arrive_time(thresholds_429);

export default function () {
  student_list_core();
}
