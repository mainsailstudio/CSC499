import { Component, OnInit } from '@angular/core';
// import { AuthenticateService } from '../authenticate/authenticate.service';
import { LoginUserService } from './login-user.service';

// export interface LoginUser {
//   email: string;
//   tempPass: string;
// }

@Component({
    selector: 'app-login',
    templateUrl: 'login.component.html',
    styleUrls: ['./login.component.css']
  })
export class LoginComponent implements OnInit {

  // register: StartRegisterUser[];

  constructor(private loginService: LoginUserService) { }

  ngOnInit() {
    // this.getHeroes();
  }

  validateRegistration(email: string, pass: string, confirmPass: string): boolean {
    if (pass === confirmPass) {
      return true;
    } else {
      return false;
    }
  }

  postData(email: string, tempPass: string): void {
    email = email.trim();
    console.log('Email is ' + email);

    this.loginService.loginUser(email).subscribe(
      suc => {
        console.log(suc);
      },
      err => {
          console.log(err );
      }
    );
  }

}

@Component({
  selector: 'app-login-success',
  templateUrl: 'login-success.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginSuccessComponent implements OnInit {

  constructor() {

  }

  ngOnInit() {

  }
}
