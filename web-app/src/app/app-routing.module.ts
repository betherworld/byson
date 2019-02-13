import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BisonDashboardComponent } from "./bison-dashboard/bison-dashboard.component";
import { RangerDashboardComponent } from "./ranger-dashboard/ranger-dashboard.component";
import { RangerVotingComponent } from "./ranger-voting/ranger-voting.component";


const routes: Routes = [
  { path: 'bison', component: BisonDashboardComponent },

  { path: 'ranger', component: RangerDashboardComponent },
  { path: 'community-vote', component: RangerVotingComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
