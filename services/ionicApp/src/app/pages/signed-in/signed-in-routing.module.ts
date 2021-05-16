import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { SignedInPage } from './signed-in.page';

const routes: Routes = [
  {
    path: '',
    component: SignedInPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class SignedInPageRoutingModule {}
