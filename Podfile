use_frameworks!
workspace 'bnwie.xcworkspace'
target "ios" do
	project 'ios/ios.xcodeporj'
	pod 'FBSDKCoreKit', '4.13.1'
	pod 'FBSDKShareKit', '4.13.1'
	pod 'FBSDKLoginKit', '4.13.1'
	pod 'Alamofire',    :git => 'https://github.com/Alamofire/Alamofire.git',    :branch => 'swift3'
end

post_install do |installer|
    installer.pods_project.targets.each do |target|
        target.build_configurations.each do |config|
            config.build_settings['ALWAYS_EMBED_SWIFT_STANDARD_LIBRARIES'] = 'NO'
        end
    end
end
