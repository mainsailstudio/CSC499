import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { RegisterUser } from './register-start.interface';
import { RegisterStartService } from './register-start.service';

@Component({
  selector: 'app-register-start',
  templateUrl: './register-start.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterStartComponent implements OnInit {

  @Output() messageEvent = new EventEmitter<string>();

  register: RegisterUser[];
 // editHero: Hero; // the hero currently being edited

  constructor(private registerService: RegisterStartService) { }

  ngOnInit() {
    // this.getHeroes();
  }

  // getHeroes(): void {
  //   this.registerService.getHeroes()
  //     .subscribe(register => this.register = register);
  // }

  add(email: string, password: string): void {
    this.register = undefined;
    email = email.trim();
    if (!email) { return; }
    password = password.trim();
    if (!email) { return; }

    // The server will generate the id for this new hero
    const newUser: RegisterUser = { email, password } as RegisterUser;
    this.registerService.addUser(newUser)
      .subscribe(register => this.register.push(register));
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

  sendEmail() {
    this.messageEvent.emit('test');
  }

}
