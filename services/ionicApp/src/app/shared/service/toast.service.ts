import {Injectable} from '@angular/core';
import {ToastController} from "@ionic/angular";

@Injectable({
  providedIn: 'root'
})
export class ToastService {

  constructor(
    private toast: ToastController
  ) {
  }

  async presentToast(text: string, time: number) {
    const toast = await this.toast.create({
      message: text,
      duration: time,
      position: 'bottom'
    });

    toast.onDidDismiss().then(() => {
      console.log('Dismissed toast');
    });

    return toast.present();
  }
}
