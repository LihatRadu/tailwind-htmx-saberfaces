package sabers

type Saber struct {
  ID int
  Name string
  Class string
  Level int
  Description string
  Image string
  Stats string
  NoblePhantasm string
}

var sabers []Saber

func Saberfaces() {
  sabers = []Saber{
    Saber{
      ID: 1,
      Name: "Artoria Pendragon",
      Class: "Saber",
      Level: 1,
      Description: "Saber Description",
      Image: "../assets/artoria.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Saber Noble Phantasm",
    },
    Saber{ 
      ID: 2,
      Name: "Okita Souji",
      Class: "Saber",
      Level: 2,
      Description: "Okita Souji Description",
      Image: "../assets/okita.jpg",
      Stats: "Saber Stats",
      NoblePhantasm: "Okita Noble Phantasm",
    },
    Saber{ 
      ID: 3,
      Name: "Nero Claudius",
      Class: "Saber",
      Level: 2,
      Description: "Nero Description",
      Image: "../assets/nero.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Nero Noble Phantasm",
    },
    Saber{ 
      ID: 4,
      Name: "Yamato Takeru",
      Class: "Saber",
      Level: 2,
      Description: "Yamato Takeru Description",
      Image: "../assets/takeru.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Takeru Noble Phantasm",
    },
    Saber{ 
      ID: 5,
      Name: "Mordred",
      Class: "Saber",
      Level: 2,
      Description: "Mordred Description",
      Image: "../assets/mordred.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Mordred Noble Phantasm",
    },
    Saber{ 
      ID: 6,
      Name: "Morgan la Faye",
      Class: "Berserker",
      Level: 2,
      Description: "Morgan la Faye Description",
      Image: "../assets/morgan.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Morgan Noble Phantasm",
    },
    Saber{ 
      ID: 7,
      Name: "Heroine X",
      Class: "Assassin",
      Level: 2,
      Description: "Heroine X Description",
      Image: "../assets/MHX1.png",
      Stats: "Saber Stats",
      NoblePhantasm: "X Noble Phantasm",
    },
    Saber{ 
      ID: 8,
      Name: "Artoria Alter(Salter)",
      Class: "Saber",
      Level: 2,
      Description: "Salter Description",
      Image: "../assets/saber_alter.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Excalibur Noble Phantasm",
    },
  }
}

func GetSabers() []Saber {
  return sabers
}
