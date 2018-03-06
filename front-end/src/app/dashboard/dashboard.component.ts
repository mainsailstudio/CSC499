import { Component, OnInit } from '@angular/core';
import { RegisterUserService } from '../register/register-user.service';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { AccountInitService } from './account-init.service';

export interface InitUser {
  id: string;
  email: string;
  fname: string;
  lname: string;
  securityLv: string;
}

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  tokenJSON = localStorage.getItem('currentUser');
  loginState = (JSON.parse(this.tokenJSON).loginState);
  // loginStateJson = localStorage.getItem('currentUser');
  constructor() { }

  ngOnInit() {
    console.log('Token JSON is ' + this.tokenJSON);
    console.log('Login state is ' + this.loginState);
  }

}

@Component({
  selector: 'app-dashboard-nav',
  templateUrl: './dashboard-nav.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardNavComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-dashboard-sidebar',
  templateUrl: './dashboard-sidebar.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardSidebarComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-dashboard-main',
  templateUrl: './dashboard-main.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardMainComponent implements OnInit {

  userDataJSON = localStorage.getItem('currentUser');
  userData = JSON.parse(this.userDataJSON);
  email = this.userData['email'];

  constructor() { }

  ngOnInit() {

  }

}

@Component({
  selector: 'app-dashboard-init',
  animations: [
    trigger(
      'myAnimation',
      [
        transition(
        ':enter', [
          style({opacity: 0}),
          animate('300ms', style({'opacity': 1}))
        ]
      ),
      // transition(
      //   ':leave', [
      //     style({'opacity': 1}),
      //     animate('100ms', style({'opacity': 0})),
      //   ]
      // )
    ]
    )
  ],
  templateUrl: './dashboard-init.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardInitComponent implements OnInit {

  userDataJSON = localStorage.getItem('currentUser');
  userData = JSON.parse(this.userDataJSON);
  email = this.userData['email'];
  userID = this.userData['id'];
  startForm = true;

  constructor(private registrationService: AccountInitService) { }

  ngOnInit() { }

  contRegistration(contRegistrationForm: any) {
    const email = this.email;
    const id = this.userID;
    console.log('User email is ' + email);

    const fname = contRegistrationForm.fname;
    console.log('User fname is ' + fname);

    const lname = contRegistrationForm.lname;
    console.log('User lname is ' + lname);

    const securityLv = contRegistrationForm.securityLv;
    console.log('User securityLv is ' + securityLv);

    this.startForm = false;
    const registerUser: InitUser = { id, email, fname, lname, securityLv } as InitUser;
    this.registrationService.initAccount(registerUser).subscribe(
      suc => {
        console.log(suc);
      },
      err => {
        console.log(err );
      }
    );
  }

}
