package sabers

type Saber struct {
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
      Name: "Artoria Pendragon",
      Class: "Saber",
      Level: 1,
      Description: "Saber Description",
      Image: "../assets/artoria.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Saber Noble Phantasm",
    },
    Saber{ 
      Name: "Okita Souji",
      Class: "Saber",
      Level: 2,
      Description: "Okita Souji Description",
      Image: "../assets/okita.jpg",
      Stats: "Saber Stats",
      NoblePhantasm: "Okita Noble Phantasm",
    },
    Saber{ 
      Name: "Nero Claudius",
      Class: "Saber",
      Level: 2,
      Description: "Nero Description",
      Image: "../assets/nero.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Nero Noble Phantasm",
    },
    Saber{ 
      Name: "Yamato Takeru",
      Class: "Saber",
      Level: 2,
      Description: "Yamato Takeru Description",
      Image: "../assets/takeru.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Takeru Noble Phantasm",
    },
    Saber{ 
      Name: "Mordred",
      Class: "Saber",
      Level: 2,
      Description: "Mordred Description",
      Image: "../assets/mordred.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Mordred Noble Phantasm",
    },
    Saber{ 
      Name: "Morgan la Faye",
      Class: "Berserker",
      Level: 2,
      Description: "Morgan la Faye Description",
      Image: "../assets/morgan.png",
      Stats: "Saber Stats",
      NoblePhantasm: "Morgan Noble Phantasm",
    },
    Saber{ 
      Name: "Heroine X",
      Class: "Assassin",
      Level: 2,
      Description: "Heroine X Description",
      Image: "../assets/MHX1.png",
      Stats: "Saber Stats",
      NoblePhantasm: "X Noble Phantasm",
    },
    Saber{ 
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
