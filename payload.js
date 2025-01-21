export default class payload {
  static getCreatePayload() {
    const name = ["Jack ", "Susan ", "Thomas ", "Ali ", "Florence "][
      Math.floor(Math.random() * 5)
    ];
    const surName = ["Dawson", "Storm", "Elva", "Abhams", "Shinya"][
      Math.floor(Math.random() * 5)
    ];

    const payload = JSON.stringify({
      name: name + surName,
      age: Math.floor(Math.random() * (18 - 10 + 1)) + 10,
      grade: ["A", "B", "C", "D"][Math.floor(Math.random() * 4)],
    });

    return payload;
  }

  static getEditPayload(id = 1) {
    const name = ["Jack ", "Susan ", "Thomas ", "Ali ", "Florence "][
      Math.floor(Math.random() * 5)
    ];
    const surName = ["Dawson", "Storm", "Elva", "Abhams", "Shinya"][
      Math.floor(Math.random() * 5)
    ];

    const payload = JSON.stringify({
      id: id,
      name: name + surName,
      age: Math.floor(Math.random() * (18 - 10 + 1)) + 10,
      grade: ["A", "B", "C", "D"][Math.floor(Math.random() * 4)],
      attendence: [true, false][Math.floor(Math.random() * 2)],
    });

    return payload;
  }
}
