import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { SignedInPageRoutingModule } from './signed-in-routing.module';

import { SignedInPage } from './signed-in.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    SignedInPageRoutingModule
  ],
  declarations: [SignedInPage]
})
export class SignedInPageModule {}
