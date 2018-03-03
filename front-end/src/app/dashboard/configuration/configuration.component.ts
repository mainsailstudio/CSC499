import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-configuration',
  templateUrl: 'configuration.component.html',
  styleUrls: ['./configuration.component.css']
})

export class ConfigurationComponent implements OnInit {

  // locks = [
  //   new LockKey(1, 'Windstorm'),
  //   new LockKey(2, 'Bombasto'),
  //   new LockKey(3, 'Magneta'),
  //   new LockKey(4, 'Tornado')
  // ];

  constructor() { }

  ngOnInit() {

  }

  postData(locks: string, keys: string): void {
    console.log('would finish things up here');
    // this.register = undefined;
    // email = email.trim();
    // console.log('Email is ' + email);
    // if (!email) { return; }
    // tempPass = tempPass.trim();
    // console.log('tempPass is ' + tempPass);
    // if (!tempPass) { return; }

    // // The server will generate the id for this new hero
    // const newUser: FinalRegisterUser = {  } as FinalRegisterUser;
    // this.registerService.finalUser(newUser).subscribe(
    //   suc => {
    //     console.log(suc);
    //   },
    //   err => {
    //       console.log(err );
    //   }
    // );
  }

}

