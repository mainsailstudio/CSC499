import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { RegisterUserService } from '../register/register-user.service';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { InitAccountService } from './init-account.service';
import { PracticeComponent } from './practice/practice.component';
import { UserConstantsService } from './user-constants/user-constants.service';

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

  userDataJSON = localStorage.getItem('currentUser');
  jwToken = (JSON.parse(this.userDataJSON).token);
  userData = JSON.parse(this.userDataJSON);
  email = this.userData['email'];
  userID = this.userData['id'];
  loginState = this.userData['loginState'];
  testLevel = this.userData['testLevel'];
  init = this.userData['init'];
  startForm = true;

  // the variable that swaps between main components
  mainActiveComponentNum = 0;

  constructor(private userConstants: UserConstantsService) { }

  ngOnInit() {
    if (this.init === false) {
      this.mainActiveComponentNum = 1;
    } else {
      this.mainActiveComponentNum = 2;
    }
  }

  // the function that swaps between the components as needed
  // kind of a view factory I guess
  swapDashboardComponent(componentNum: number, delay: number) {
    Observable.timer(delay)
        .subscribe(i => {
          this.mainActiveComponentNum = componentNum;
        });
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

  init: boolean;
  testLevel: string;
  userDataJSON = localStorage.getItem('currentUser');
  userData = JSON.parse(this.userDataJSON);

  constructor() { }

  ngOnInit() {
    this.userDataJSON = localStorage.getItem('currentUser');
    this.userData = JSON.parse(this.userDataJSON);
    this.init = this.userData['init'];
    this.testLevel = this.userData['testLevel'];
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
  jwToken = (JSON.parse(this.userDataJSON).token);
  userData = JSON.parse(this.userDataJSON);
  email = this.userData['email'];
  userID = this.userData['id'];
  loginState = this.userData['loginState'];
  testLevel = this.userData['testLevel'];
  init = this.userData['init'];
  startForm = true;

  constructor(private initAccountService: InitAccountService) { }

  ngOnInit() { }

  contRegistration(contRegistrationForm: any) {
    const email = this.email;
    const id = this.userID;

    const fname = contRegistrationForm.fname;

    const lname = contRegistrationForm.lname;

    const securityLv = contRegistrationForm.securityLv;

    this.startForm = false;
    const registerUser: InitUser = { id, email, fname, lname, securityLv } as InitUser;
    this.initAccountService.initAccount(registerUser, this.jwToken).subscribe(
      suc => {
      },
      err => {
      }
    );
  }

}


@Component({
  selector: 'app-dashboard-practice',
  templateUrl: './dashboard-practice.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardPracticeComponent implements OnInit {

  constructor() { }

  ngOnInit() {

  }

}

@Component({
  selector: 'app-dashboard-hints',
  templateUrl: './dashboard-hints.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardHintsComponent implements OnInit {

  constructor() { }

  ngOnInit() {

  }

}

@Component({
  selector: 'app-dashboard-about',
  templateUrl: './dashboard-about.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardAboutComponent implements OnInit {

  constructor() { }

  ngOnInit() {

  }

}
