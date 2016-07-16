//
//  SecondViewController.swift
//  bnwie
//
//  Created by TruongSinh Tran-Nguyen on 7/10/16.
//  Copyright © 2016 TruongSinh Tran-Nguyen. All rights reserved.
//

import UIKit
import FBSDKLoginKit
class LoginVC: UIViewController {
    private let readPermissions = ["public_profile", "email", "user_friends"]
    @IBOutlet weak var loginButton: UIButton!
    override func viewDidLoad() {
        super.viewDidLoad()
    }

    override func viewDidAppear(_ animated: Bool) {
        if let accessToken = FBSDKAccessToken.currentAccessToken() {
            loginButton.hidden = true
            performSegueWithIdentifier("Landing2Main", sender: self)
        }
    }

    @IBAction func doSignInWithFb(_ sender: UIButton, forEvent event: UIEvent) {
        FBSDKLoginManager()
            .logInWithReadPermissions(readPermissions, handler: handleLogin)
    }
    private func handleLogin (_ result: FBSDKLoginManagerLoginResult?, _ err: NSError?) {
        print("FBLOGIN private func handleLogin")
            if let err = err {
            return handleLoginError(err)

        }
            guard let result = result else {
            return handleLoginUnexpected()
        }
        if result.isCancelled == true {
            return handleLoginCancel()
        }
        return handleLoginSuccess(result)
    }
    private func handleLoginError(_ err: NSError) {
        print("FBLOGIN err")
        print("FBLOGIN Error writing to URL: \(err)")
        NSLog("FBLOGIN %@",err)
    }

    private func handleLoginUnexpected() {
        print("FBLOGINunex")
    }

    private func handleLoginCancel() {
        print("FBLOGINcanc")
    }

    private func handleLoginSuccess(_ result: FBSDKLoginManagerLoginResult) {
        print("FBLOGINsuccess")
    }

}
