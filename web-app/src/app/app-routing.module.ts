import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BisonDashboardComponent } from "./bison-dashboard/bison-dashboard.component";
import { RangerDashboardComponent } from "./ranger-dashboard/ranger-dashboard.component";
import { RangerVotingComponent } from "./ranger-voting/ranger-voting.component";
import { DonorImpactDashboardComponent } from "./donor-impact-dashboard/donor-impact-dashboard.component";
import { DonorHomeComponent } from "./donor-home/donor-home.component";
import { RangerProofComponent } from "./ranger-proof/ranger-proof.component";
import { DonorEnvironmentComponent } from "./donor-environment/donor-environment.component";
import { WildlifeDashboardComponent } from "./wildlife-dashboard/wildlife-dashboard.component";


const routes: Routes = [
  { path: 'donor/home', component: DonorHomeComponent },
  { path: 'donor/environment', component: DonorEnvironmentComponent },

  { path: 'bison/:bison_id', component: BisonDashboardComponent },
  { path: 'wildlife/:wild_id', component: WildlifeDashboardComponent },

  { path: 'ranger/home', component: RangerDashboardComponent },
  { path: 'ranger/community-vote', component: RangerVotingComponent },
  { path: 'ranger/impact', component: DonorImpactDashboardComponent },
  { path: 'ranger/proof', component: RangerProofComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
