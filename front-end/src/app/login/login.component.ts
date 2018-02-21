import { LoginService } from './login.service';
import { Component } from '@angular/core';

@Component({
    selector: 'app-login',
    templateUrl: 'login.component.html',
    styleUrls: ['./login.component.css']
  })
  export class LoginComponent {
    login;
    constructor(service: LoginService) {
      this.login = service.login();
    }
   }
