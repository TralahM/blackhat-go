#!/usr/bin/env bash

find -type f -name "*.go" -exec sed -i 's/blackhat-go\/bhg/TralahM\/blackhat-go/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-2/chapter2_TCP_and_Go_Scanners_and_Proxies/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-3/chapter3_HTTP_Clients_Remote_Interaction_with_Tools/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-4/chapter4_HTTP_Servers_Routing_and_Middleware/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-5/chapter5_Exploiting_DNS_Recon_and_More/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-6/chapter6_SMB_and_NTLM_A_Peek_Down_the_Rabbit_Hole/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-7/chapter7_Databases_and_Filesystems_Pilfering_and_Abusing/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-8/chapter8_Packet_Processing_Living_on_the_Wire/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-9/chapter9_Exploit_Code_Writing_and_Porting/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-10/chapter10_Extendable_Tools_Using_Go_Plugins_and_LUA/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-11/chapter11_Cryptography_Implementing_and_Attacking/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-12/chapter12_Windows_System_Interaction_and_Analysis/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-13/chapter13_Steganography_Hiding_Data/g' {} \;
find -type f -name "*.go" -exec sed -i 's/ch-14/chapter14_Command_and_Control_Building_a_RAT/g' {} \;

find -type f -name "*.go" -exec sed -i 's/TralahM/tralahm/g' {} \;
find -type f -name "*.go" -exec sed -i 's/github.com\/bhg/github.com\/tralahm\/blackhat-go/g' {} \;

git status
