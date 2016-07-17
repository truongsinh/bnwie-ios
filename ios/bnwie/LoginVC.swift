//
//  SecondViewController.swift
//  bnwie
//
//  Created by TruongSinh Tran-Nguyen on 7/10/16.
//  Copyright © 2016 TruongSinh Tran-Nguyen. All rights reserved.
//

import UIKit
import FBSDKLoginKit
import Alamofire
let endpoint = "http://sugar.smarpsocial.com:8080/api"
//let endpoint = "https://www.google.com"
class LoginVC: UIViewController {
    private let readPermissions = ["public_profile", "email", "user_friends"]
    private var profileFromResponseJson: AnyObject?
    @IBOutlet weak var loginButton: UIButton!
    @IBOutlet weak var loginIndicator: UIActivityIndicatorView!
    override func viewDidLoad() {
        super.viewDidLoad()
        loginIndicator.stopAnimating()
    }

    override func viewDidAppear(_ animated: Bool) {
        if loginButton.hidden == false && FBSDKAccessToken.currentAccessToken() != nil {
            loginButton.hidden = true
            loginIndicator.startAnimating()
            handleLoginSuccess(FBSDKAccessToken.currentAccessToken())
        }
    }

    @IBAction func doSignInWithFb( sender: UIButton, forEvent event: UIEvent) {
        loginButton.hidden = true
        loginIndicator.startAnimating()
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
        return handleLoginSuccess(result.token)
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

    override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
        if (segue.identifier == "Landing2Main") {
            
            guard let avatarUrl = profileFromResponseJson?["AvatarUrl"] as? String else {
                return
            }
            
            let barViewControllers = segue.destinationViewController as! UITabBarController
            let destinationViewController = barViewControllers.viewControllers![0] as! MyProfileViewController
            destinationViewController.preapreImage(avatarUrl)
        }
    }
    private func handleLoginSuccess(_ token: FBSDKAccessToken) {
        Alamofire
            .request(
                .POST,
                endpoint + "/user/authenticate",
                parameters: [
                    "SocnetType": "facebook",
                    "SocnetId": token.userID,
                    "SocnetToken": token.tokenString,
                ],
                encoding: .JSON
            )
            .responseJSON { response in
                print("request")
                print(response.request)  // original URL request
                print("response")
                print(response.response) // URL response
                print("data")
                print(response.data)     // server data
                print("result")
                print(response.result.debugDescription)   // result of response serialization
                print(response.result.description)   // result of response serialization
                print(response.result.error)   // result of response serialization
                print("done")
                
                if let JSON = response.result.value {
                    self.profileFromResponseJson = JSON
                    self.performSegueWithIdentifier("Landing2Main", sender: self)
                    
                }
        }
        print("FBLOGINsuccess")
    }

}
