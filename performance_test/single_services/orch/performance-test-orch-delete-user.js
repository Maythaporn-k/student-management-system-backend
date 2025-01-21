import { delete_user_orch } from "../loadtest_api/delete_user_orch.js";
import Options from "../../option.config.js";

export const thresholds_409 = {
  "http_req_failed{status:409}": ["rate<0.5"],
};

export const options = Options.const(thresholds_409);

export default function () {
  delete_user_orch();
}
