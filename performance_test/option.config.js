export default class Options {
  static const(
    vus = 10, //default
    duration = "60s", //default
    thresholds = null //default
  ) {
    const options = {
      vus: vus, // จำนวนผู้ใช้งานเริ่มต้น
      duration: duration, // ระยะเวลารันการทดสอบ
    };

    if (thresholds != null) {
      options.thresholds = thresholds;
    }

    return options;
  }

  static const_arrive_time(
    rate = 5,
    timeUnit = "10s",
    duration = "30s",
    preAllocatedVUs = 20,
    thresholds = null
  ) {
    const options = {
      scenarios: {
        student_loadtest: {
          executor: "constant-arrival-rate",
          rate: rate, // 5 requests ต่อ timeUnit
          timeUnit: timeUnit, // ต่อ 1 วินาที
          duration: duration, // ทดสอบ 1 นาที
          preAllocatedVUs: preAllocatedVUs, // เตรียม VUs ไว้ 20
        },
      },
    };

    // หาก thresholds ไม่เป็น null ให้เพิ่มเข้าไปใน options
    if (thresholds != null) {
      options.thresholds = thresholds;
    }

    return options; // คืนค่ากลับ
  }
}
