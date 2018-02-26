import { Component, OnInit, Input } from '@angular/core';
import { trigger, state, transition, style, animate } from '@angular/animations';

import { RegisterUser } from './register-start.interface';
import { RegisterStartService } from './register-start.service';
import { LockKey } from './lock-key/lock-key';

@Component({
    selector: 'app-register',
    templateUrl: 'register.component.html',
    styleUrls: ['./register.component.css']
  })
export class RegisterComponent implements OnInit {

  @Input() email: string;

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-register-start',
  templateUrl: 'register-start.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterStartComponent implements OnInit {

  register: RegisterUser[];

  constructor(private registerService: RegisterStartService) { }

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

  add(email: string, tempPass: string): void {
    this.register = undefined;
    email = email.trim();
    console.log('Email is ' + email);
    if (!email) { return; }
    tempPass = tempPass.trim();
    console.log('tempPass is ' + tempPass);
    if (!tempPass) { return; }

    // The server will generate the id for this new hero
    const newUser: RegisterUser = { email, tempPass } as RegisterUser;
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

  register: RegisterUser[];

  constructor(private registerService: RegisterStartService) { }

  ngOnInit() {

  }

  add(email: string, tempPass: string): void {
    this.register = undefined;
    email = email.trim();
    console.log('Email is ' + email);
    if (!email) { return; }
    tempPass = tempPass.trim();
    console.log('tempPass is ' + tempPass);
    if (!tempPass) { return; }

    // The server will generate the id for this new hero
    const newUser: RegisterUser = { email, tempPass } as RegisterUser;
    this.registerService.addUser(newUser).subscribe(
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
  selector: 'app-register-final',
  templateUrl: 'register-final.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterFinalComponent implements OnInit {

  locks = [
    new LockKey(1, 'Windstorm'),
    new LockKey(2, 'Bombasto'),
    new LockKey(3, 'Magneta'),
    new LockKey(4, 'Tornado')
  ];

  register: RegisterUser[];

  constructor(private registerService: RegisterStartService) { }

  ngOnInit() {

  }

  get(email: string, tempPass: string): void {
    this.register = undefined;
    email = email.trim();
    console.log('Email is ' + email);
    if (!email) { return; }
    tempPass = tempPass.trim();
    console.log('tempPass is ' + tempPass);
    if (!tempPass) { return; }

    // The server will generate the id for this new hero
    const newUser: RegisterUser = { email, tempPass } as RegisterUser;
    this.registerService.addUser(newUser).subscribe(
      suc => {
        console.log(suc);
      },
      err => {
          console.log(err );
      }
    );
  }

}
