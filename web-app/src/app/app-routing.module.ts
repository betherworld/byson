import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BisonDashboardComponent } from "./bison-dashboard/bison-dashboard.component";
import { RangerDashboardComponent } from "./ranger-dashboard/ranger-dashboard.component";
import { RangerVotingComponent } from "./ranger-voting/ranger-voting.component";
import { DonorImpactDashboardComponent } from "./donor-impact-dashboard/donor-impact-dashboard.component";
import {DonorHomeComponent} from "./donor-home/donor-home.component";


const routes: Routes = [
  { path: 'donor/home', component: DonorHomeComponent },
  { path: 'donor/bison', component: BisonDashboardComponent },
  { path: 'donor/impact', component: DonorImpactDashboardComponent },

  { path: 'ranger/home', component: RangerDashboardComponent },
  { path: 'ranger/community-vote', component: RangerVotingComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
