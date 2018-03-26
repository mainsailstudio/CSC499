import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterStartComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { LogoutComponent } from './login/logout.component';
import { LandingComponent } from './landing/landing.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { DashboardComponent,
         DashboardPracticeComponent,
         DashboardHintsComponent,
         DashboardAboutComponent } from './dashboard/dashboard.component';

// test imports
import { LandingTestComponent } from './landing/landing-test.component';
import { DashboardTestComponent } from './dashboard/dashboard-test.component';

// new imports for authentication
import { LoginNewComponent } from './login/login-new.component';
import { AuthGuard } from './_auth-guard/auth.guard';
import { PracticeComponent } from './dashboard/practice/practice.component';
import { AboutComponent } from './about/about.component';
import { HintsComponent } from './hints/hints.component';
import { UsabilityTestComponent } from './dashboard/usability-test/usability-test.component';

const routes: Routes = [
  { path: '',  component: LandingTestComponent },
  // { path: 'test',  component: LandingComponent },
  // { path: 'register', component: RegisterStartComponent },
  // { path: 'login', component: LoginComponent },
  { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
  { path: 'practice',  component: DashboardPracticeComponent, canActivate: [AuthGuard] },
  { path: 'logout', component: LogoutComponent },
  { path: 'about', component: DashboardAboutComponent },
  { path: 'hints', component: DashboardHintsComponent },
  { path: 'instructions', component: UsabilityTestComponent },

  // { path: 'dashboard-test', component: DashboardTestComponent, canActivate: [AuthGuard] },
  { path: '**', component: NotFoundComponent } // make sure this is always at the bottom so it doesn't superscede legitimate routes
];

@NgModule({
  imports: [
     RouterModule.forRoot(routes, {useHash: true})
    ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }

