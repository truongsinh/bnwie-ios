//
//  ToggleTableViewCell.swift
//  ios
//
//  Created by TruongSinh Tran-Nguyen on 7/17/16.
//  Copyright © 2016 TruongSinh Tran-Nguyen. All rights reserved.
//

import UIKit


struct ServiceStruct {
    var emoji: String
    var value: String
    var text: String
}

class ToggleTableViewCell: UITableViewCell {
    init(service: ServiceStruct,on: Bool) {
        super.init(style: UITableViewCellStyle.Default, reuseIdentifier: nil)
        self.textLabel?.text = service.emoji + " " + service.text
        let s = UISwitch()
        s.on = on
        s.addTarget(
            self, action: #selector(ToggleTableViewCell.switchIsChanged(_:)),
            forControlEvents: UIControlEvents.ValueChanged
        )
        self.accessoryView = s
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    func switchIsChanged(mySwitch: UISwitch) {
        if mySwitch.on {
//            switchState.text = "UISwitch is ON"
        } else {
//            switchState.text = "UISwitch is OFF"
        }
        print("switch state \(mySwitch.on)")
    }
    
}
