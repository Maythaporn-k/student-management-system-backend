import { delete_user_orch } from "../../loadtest_api/orch/delete_user_orch.js";
import Options from "../../option.config.js";

export const thresholds = {
  "http_req_failed{status:405}": ["rate>0.05"],
  "http_req_failed{status:409}": ["rate>0.05"],
};

export const options = Options.const(thresholds);

export default function () {
  delete_user_orch();
}
