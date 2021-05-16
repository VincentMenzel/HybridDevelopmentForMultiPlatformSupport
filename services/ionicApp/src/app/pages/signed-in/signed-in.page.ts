import {Component, OnInit} from '@angular/core';
import {AuthService} from '../../shared/service/auth.service';
import {IUser} from '../../shared/interface/user';
import {NavController} from '@ionic/angular';

@Component({
  selector: 'app-signed-in',
  templateUrl: './signed-in.page.html',
  styleUrls: ['./signed-in.page.scss'],
})
export class SignedInPage implements OnInit {

  public user: null | IUser;

  constructor(private authService: AuthService, private navController: NavController) {
  }


  ngOnInit() {
    this.authService.onUserChange.subscribe(user => {
      this.user = user;
      if (this.user === null) {
        this.navController.navigateRoot('sign-in');
      }
    });
  }

  onBack($event: MouseEvent) {
    this.navController.back();
  }
}
