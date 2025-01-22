import { student_list_orch } from "../../loadtest_api/orch/student_list_orch.js";
import Options from "../../option.config.js";

export const thresholds = {
  "http_req_failed{status:405}": ["rate>0.05"],
  "http_req_failed{status:429}": ["rate>0.05"],
};

export const options = Options.const_arrive_time(thresholds);

export default function () {
  student_list_orch();
}
