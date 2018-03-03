import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { trigger, state, transition, style, animate } from '@angular/animations';

import { RegisterUserService } from './register-user.service';
import { LockKey } from './lock-key/lock-key';

export interface StartRegisterUser {
  id: number;
  email: string;
  tempPass: string;
}

@Component({
  selector: 'app-register-start',
  templateUrl: 'register-start.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterStartComponent implements OnInit {

  register: StartRegisterUser[];

  constructor(private registerService: RegisterUserService) { }

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
    this.register = undefined;
    email = email.trim();
    console.log('Email is ' + email);
    if (!email) { return; }
    tempPass = tempPass.trim();
    console.log('tempPass is ' + tempPass);
    if (!tempPass) { return; }

    // The server will generate the id for this new hero
    const newUser: StartRegisterUser = { email, tempPass } as StartRegisterUser;
    console.log('newUser to be posted is ' + newUser);
    this.registerService.addUser(newUser).subscribe(
      suc => {
        console.log(suc);
      },
      err => {
          console.log(err );
      }
    );
  }

  // delete(hero: Hero): void {
  //   this.register = this.register.filter(h => h !== hero);
  //   this.registerService.deleteHero(hero.id).subscribe();
  //   /*
  //   // oops ... subscribe() is missing so nothing happens
  //   this.registerService.deleteHero(hero.id);
  //   */
  // }

  // edit(hero) {
  //   this.editHero = hero;
  // }

  // search(searchTerm: string) {
  //   this.editHero = undefined;
  //   if (searchTerm) {
  //     this.registerService.searchHeroes(searchTerm)
  //       .subscribe(register => this.register = register);
  //   }
  // }

  // update() {
  //   if (this.editHero) {
  //     this.registerService.updateHero(this.editHero)
  //       .subscribe(hero => {
  //         // replace the hero in the register list with update from server
  //         const ix = hero ? this.register.findIndex(h => h.id === hero.id) : -1;
  //         if (ix > -1) { this.register[ix] = hero; }
  //       });
  //     this.editHero = undefined;
  //   }
  // }
}

@Component({
  selector: 'app-register-continue',
  templateUrl: 'register-continue.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterContinueComponent implements OnInit {

  constructor() { }

  ngOnInit() {

  }

  // postData(fname: string, lname: string, securityLv: number): void {
  //   this.register = undefined;
  //   fname = fname.trim();
  //   console.log('fname is ' + fname);
  //   if (!fname) { return; }

  //   lname = lname.trim();
  //   console.log('lname is ' + lname);
  //   if (!lname) { return; }

  //   console.log('securityLv is ' + securityLv);
  //   if (!lname) { return; }

  //   // The server will generate the id for this new hero
  //   const newUser: ContinueRegisterUser = { fname, lname, securityLv } as ContinueRegisterUser;
  //   this.registerService.continueUser(newUser).subscribe(
  //     suc => {
  //       console.log(suc);
  //     },
  //     err => {
  //         console.log(err );
  //     }
  //   );
  // }
}
