<h1>gorestledger</h1>
<p><a href="https://www.hyperledger.org/projects/fabric"><img src="https://www.hyperledger.org/wp-content/uploads/2016/09/logo_hl_new.png" alt="N|Solid"></a></p>
<p>GoRestLedger is a REST based API implementation written in Go language to demonstrate the Hyperledger Fabric Blockchain framework. The project repo has been designed to upload employee records into the blockchain and also has the functionality to update, delete the record stored securely in the Hyperledger framework.</p>

<p> However, this explanation guide does not explain how Hyperledger Fabric works, so for the information, you can follow at <a href="https://www.hyperledger.org/projects/fabric">Hyperledger.</a> </p>

<h4><a id="Installation_6"></a>Installation</h4>
<p>Employeeledger requires <a href="https://www.docker.com/">Docker</a> &amp; <a href="https://golang.org/">Go</a> to run.</p>
<h3><a id="Docker_10"></a>Docker</h3>
<pre><code class="language-sh">$ sudo apt install docker.io
$ sudo apt install docker-compose
</code></pre>
<h2><a id="Go_15"></a>Go</h2>
<h4><a id="Installation_16"></a>Installation</h4>
<pre><code class="language-sh">$ sudo apt-get update
$ sudo apt-get install golang-go
</code></pre>
<h4><a id="Set_your_Go_path_as_environmental_variable_21"></a>Set your Go path as environmental variable</h4>
<h6><a id="add_these_following_variable_into_the_profile_22"></a>add these following variable into the profile</h6>
<pre><code class="language-sh">$ <span class="hljs-built_in">export</span> GOPATH=<span class="hljs-variable">$HOME</span>/go
$ <span class="hljs-built_in">export</span> PATH=<span class="hljs-variable">$PATH</span>:/usr/<span class="hljs-built_in">local</span>/go/bin:<span class="hljs-variable">$GOPATH</span>/bin
</code></pre>
<h6><a id="then_27"></a>then</h6>
<pre><code class="language-sh">$ <span class="hljs-built_in">source</span> ~/.profile
$ go version
$ go version go1.<span class="hljs-number">11</span> linux/amd64
</code></pre>
<h3><a id="Build_Your_Network_34"></a>Build Your Network</h3>
<p>This sample Hyperledger Fabric blockchain network is built on a single organization consisting of two peer nodes. There are few prerequisites to follow to set up and install a blockchain network in the docker container.</p>
<h5><a id="Prerequisites_38"></a>Prerequisites</h5>
<ul>
<li>There are few binaries needs to be download to create the network.</li>
<li>These binaries contains various commands to setup , install and execute contains written in .yaml file.</li>
<li>Command Reference Guide : <a href="https://hyperledger-fabric.readthedocs.io/en/release-1.4/command_ref.html">https://hyperledger-fabric.readthedocs.io/en/release-1.4/command_ref.html</a></li>
</ul>
<h6><a id="You_can_choose_any_of_the_following_link_based_on_you_operating_system_and_hardware_architecture_of_your_system_43"></a>You can choose any of the following link based on you operating system and hardware architecture of your system.</h6>
<table class="table table-striped table-bordered">
<thead>
<tr>
<th>Arch</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/linux-amd64-1.0.5/hyperledger-fabric-linux-amd64-1.0.5.tar.gz">Linux AMD 64</a></td>
</tr>
<tr>
<td><a href="https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/linux-s390x-1.0.5/hyperledger-fabric-linux-s390x-1.0.5.tar.gz">Linux s390x</a></td>
</tr>
<tr>
<td><a href="https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/linux-ppc64le-1.0.5/hyperledger-fabric-linux-ppc64le-1.0.5.tar.gz">Linux PPC64le</a></td>
</tr>
<tr>
<td><a href="https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/windows-amd64-1.0.5/hyperledger-fabric-windows-amd64-1.0.5.tar.gz">Windows AMD 64</a></td>
</tr>
<tr>
<td><a href="https://nexus.hyperledger.org/content/repositories/releases/org/hyperledger/fabric/hyperledger-fabric/darwin-amd64-1.0.5/hyperledger-fabric-darwin-amd64-1.0.5.tar.gz">Darwin AMD 64</a></td>
</tr>
</tbody>
</table>
<h5><a id="Configuration_53"></a>Configuration</h5>
<p>There are few configuration files is written, which composed of various configuration details of a network like Certificate Authority (CA), Peer , Orderer , CouchDB.</p>
<h5><a id="cryptoconfigyaml_57"></a>crypto-config.yaml</h5>
<p>This configuration file will generate few certificates and key for the organization and it’s related entities like peers, admin, orderer. The cryptogen binary will take the config file as input and after execution, it will create a crypto-config folder in the config directory, which will contains all the generated certificates and key.</p>
<pre><code>./bin/cryptogen generate --config=./crypto-config.yaml
</code></pre>
<h5><a id="configtxyaml_62"></a>configtx.yaml</h5>
<p>This config file will contains complete details of a channel related to an organization. It will create three artifacts for a network.</p>
<ul>
<li>
<h6><a id="orderergenesisblock__This_will_initialize_the_Fabrics_orderer_65"></a>orderer.genesis.block : This will initialize the Fabric’s orderer</h6>
<pre><code>  ./bin/configtxgen -profile GoRestLedger -outputBlock ./artifacts/orderer.genesis.block
</code></pre>
</li>
<li>
<h6><a id="channeltx__Channel_is_a_private_network_between_peers_to_communicate_in_a_network_66"></a>gorestledger.channel.tx : Channel is a private network between peers to communicate in a network</h6>
<pre><code>  ./bin/configtxgen -profile GoRestLedger -outputCreateChannelTx ./artifacts/gorestledger.channel.tx -channelID gorestledger
</code></pre>
</li>    
<li>
<h6><a id="org1employeeledgeranchorstx__This_artifact_will_allow_the_peers_to_interact_with_each_other_in_a_network_68"></a>org1.employeeledger.anchors.tx : This artifact will allow the peers to interact with each other in a network.</h6>
<pre><code>  ./bin/configtxgen -profile GoRestLedger -outputAnchorPeersUpdate ./artifacts/org1.gorestledger.anchors.tx -channelID gorestledger -asOrg GoRestLedgerOrganization1
</code></pre>
</li>
</ul>
<h6><a id="You_can_find_a_shell_script_configsh_in_the_fixtures_folder_it_will_generate_all_the_prerequisites_configurations_of_a_network_So_you_can_directly_execute_script_to_skip_all_the_manual_steps_72"></a>You can find a shell script &quot;<a href="http://config.sh">config.sh</a>&quot; in the fixtures folder, it will generate all the prerequisites configurations of a network. So, you can directly execute script to skip all the manual steps.</h6>
<h5><a id="DockerCompose_75"></a>Docker-Compose</h5>
<p>Now, we need to deploy the configuration details into a docker container, so we need to use Docker Compose. There will docker-compose.yaml configuration file, which will contain all config details for Orderer, Certificate Authority, Peer, Couch DB.</p>
<ul>
<li>
<p>The compose file can be deploy into a network, by executing following command</p>
<pre><code> docker-compose up -d // the docker-compose.yaml has to be located at same command location
</code></pre>
</li>
</ul>
<h6><a id="And_also_replace_the_newly_generated_CA_key_which_can_be_found_at_cryptoconfigpeerOrganizationsorg1employeeledgercomcasome_random_generated_sk_file_at_the_dockercomposeyaml__line_44_50__55_82"></a>And also replace the newly generated CA key, which can be found at crypto-config/peerOrganizations/org1.employee.ledger.com/ca/some random generated sk file, at the docker-compose.yaml : line 44, 50 &amp; 55</h6>
<p>So, all done , your blockchain network is now deployed into a docker container.</p>
<p>You can check by executing following command</p>
<pre><code>docker ps
</code></pre>

<h4>Dependency Issues</h4>
<ol>
   <li>
      Hyperledger fabric-sdk-go is still in development. If you do dep ensure for each <b>Gopkg.toml</b> in <b>gorestledger</b> and <b>Chaincode</b>, it will download the govendor folder for each module but it will have some compilation issues while building the project. I have corrected the error for both <b>gorestledger and Chaincode</b> folder.
   </li>
   <li>
   Please download the vendor folder and add it in your project repo.   
      
   gorestledger - https://www.dropbox.com/s/ry1jmw0y9xliose/vendor.zip?dl=0
   
   Chaincode - https://www.dropbox.com/s/31nnqflpqwaywoa/vendor.zip?dl=0
   </li>
   <li>
   <b>Add vendor folders at the location where Gopkg.toml file is located.</b>
   </li>
</ol>

<h4><a id="Run_the_application_90"></a>Run the application</h4>
<ul>
<li>As you have sucssefully deployed your blockchain network. Now you can run the application.</li>
<li>There is a master Makefile , you can find in the project root directory.</li>
<li>Just type “make” in the command line and it will take few minuetes to start the server.</li>
<li>If all goes well, you can see server running at port 4000</li>
<li>Now, you can use postman to test the APIs at <a href="http://localhost:6000">http://localhost:6000</a></li>
</ul>

<br>

<h3>REST APIs end points</h3>

<h4><a id="apiuser_register_0"></a>/api/user_register</h4>
<h4><a id="POST_1"></a>POST</h4>
<h6><a id="Request_2"></a>Request</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"name"</span>: <span class="hljs-string">"Alice McDonnell"</span>,
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
    <span class="hljs-string">"password"</span>: <span class="hljs-string">"Hello@123"</span>,
    <span class="hljs-string">"company"</span>: <span class="hljs-string">"Apple"</span>,
    <span class="hljs-string">"occupation"</span>: <span class="hljs-string">"Architect"</span>,
    <span class="hljs-string">"salary"</span>: <span class="hljs-string">"85,000"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"Admin"</span>
}
</code></pre>
<h6><a id="Response_15"></a>Response</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"id"</span>: <span class="hljs-string">"eDUwOTo6Q049YW5zaHVtYW4ucGF0dG5haWs1MjRAZ21haWwuY29tLE9VPXVzZXIrT1U9b3JnMTo6Q049Y2Eub3JnMS5nby5yZXN0LmxlZGdlci5jb20sTz1vcmcxLmdvLnJlc3QubGVkZ2VyLmNvbSxMPVNhbiBGcmFuY2lzY28sU1Q9Q2FsaWZvcm5pYSxDPVVT"</span>,
    <span class="hljs-string">"name"</span>: <span class="hljs-string">"Alice McDonnell"</span>,
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
    <span class="hljs-string">"company"</span>: <span class="hljs-string">"Apple"</span>,
    <span class="hljs-string">"occupation"</span>: <span class="hljs-string">"Architect"</span>,
    <span class="hljs-string">"salary"</span>: <span class="hljs-string">"85,000"</span>,
    <span class="hljs-string">"token"</span>: <span class="hljs-string">"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NjExNTU5OTIsInVzZXIiOiJhbnNodW1hbi5wYXR0bmFpazUyNEBnbWFpbC5jb20ifQ.dIpTAgxnh-4lw06IQwX_bT4fhpvCUrwSZ7EnGVs46EE"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"Admin"</span>
}
</code></pre>
<h4><a id="apiuser_login_30"></a>/api/user_login</h4>
<h4><a id="POST_31"></a>POST</h4>
<h6><a id="Request_32"></a>Request</h6>
<pre><code class="language-sh">{
 <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
 <span class="hljs-string">"password"</span>: <span class="hljs-string">"Hello@456"</span>
}
</code></pre>
<h6><a id="Response_39"></a>Response</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"id"</span>: <span class="hljs-string">"eDUwOTo6Q049YW5zaHVtYW4ucGF0dG5haWs1MjRAZ21haWwuY29tLE9VPXVzZXIrT1U9b3JnMTo6Q049Y2Eub3JnMS5nby5yZXN0LmxlZGdlci5jb20sTz1vcmcxLmdvLnJlc3QubGVkZ2VyLmNvbSxMPVNhbiBGcmFuY2lzY28sU1Q9Q2FsaWZvcm5pYSxDPVVT"</span>,
    <span class="hljs-string">"name"</span>: <span class="hljs-string">"Alice McDonnell"</span>,
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
    <span class="hljs-string">"company"</span>: <span class="hljs-string">"Apple"</span>,
    <span class="hljs-string">"occupation"</span>: <span class="hljs-string">"Architect"</span>,
    <span class="hljs-string">"salary"</span>: <span class="hljs-string">"85,000"</span>,
    <span class="hljs-string">"token"</span>: <span class="hljs-string">"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NjExNTU5OTIsInVzZXIiOiJhbnNodW1hbi5wYXR0bmFpazUyNEBnbWFpbC5jb20ifQ.dIpTAgxnh-4lw06IQwX_bT4fhpvCUrwSZ7EnGVs46EE"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"Admin"</span>
}
</code></pre>
<h4><a id="apiget_users_54"></a>/api/get_users</h4>
<h4><a id="GET_55"></a>GET</h4>
<h5><a id="Headers__Token__token_value_56"></a>Headers - (Token : token_value)</h5>
<h6><a id="Response_58"></a>Response</h6>
<pre><code class="language-sh">list of user records 
</code></pre>
<h4><a id="apiget_user_63"></a>/api/get_user</h4>
<h4><a id="GET_64"></a>GET</h4>
<h5><a id="Headers__Token__token_value_65"></a>Headers - (Token : token_value)</h5>
<h6><a id="Request_67"></a>Request</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>
}
</code></pre>
<h6><a id="Response_74"></a>Response</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"id"</span>: <span class="hljs-string">"eDUwOTo6Q049YW5zaHVtYW4ucGF0dG5haWs1MjRAZ21haWwuY29tLE9VPXVzZXIrT1U9b3JnMTo6Q049Y2Eub3JnMS5nby5yZXN0LmxlZGdlci5jb20sTz1vcmcxLmdvLnJlc3QubGVkZ2VyLmNvbSxMPVNhbiBGcmFuY2lzY28sU1Q9Q2FsaWZvcm5pYSxDPVVT"</span>,
    <span class="hljs-string">"name"</span>: <span class="hljs-string">"Alice McDonnell"</span>,
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
    <span class="hljs-string">"company"</span>: <span class="hljs-string">"Apple"</span>,
    <span class="hljs-string">"occupation"</span>: <span class="hljs-string">"Architect"</span>,
    <span class="hljs-string">"salary"</span>: <span class="hljs-string">"85,000"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"Admin"</span>
}
</code></pre>
<h4><a id="apiupdate_user_87"></a>/api/update_user</h4>
<h4><a id="PUT_88"></a>PUT</h4>
<h5><a id="Headers__Token__token_value_89"></a>Headers - (Token : token_value)</h5>
<h6><a id="Request_91"></a>Request</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"name"</span>: <span class="hljs-string">"Alice Wrick"</span>,
    <span class="hljs-string">"company"</span>: <span class="hljs-string">"Apple Computer"</span>,
    <span class="hljs-string">"occupation"</span>: <span class="hljs-string">"Architect Solution"</span>,
    <span class="hljs-string">"salary"</span>: <span class="hljs-string">"85,000"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"Admin"</span>
}
</code></pre>
<h6><a id="Response_101"></a>Response</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"id"</span>: <span class="hljs-string">"eDUwOTo6Q049YW5zaHVtYW4ucGF0dG5haWs1MjRAZ21haWwuY29tLE9VPXVzZXIrT1U9b3JnMTo6Q049Y2Eub3JnMS5nby5yZXN0LmxlZGdlci5jb20sTz1vcmcxLmdvLnJlc3QubGVkZ2VyLmNvbSxMPVNhbiBGcmFuY2lzY28sU1Q9Q2FsaWZvcm5pYSxDPVVT"</span>,
    <span class="hljs-string">"name"</span>: <span class="hljs-string">"Alice Wrick"</span>,
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
    <span class="hljs-string">"company"</span>: <span class="hljs-string">"Apple Computer"</span>,
    <span class="hljs-string">"occupation"</span>: <span class="hljs-string">"Architect Solution"</span>,
    <span class="hljs-string">"salary"</span>: <span class="hljs-string">"85,000"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"Admin"</span>
}
</code></pre>
<h4><a id="apidelete_user_114"></a>/api/delete_user</h4>
<h4><a id="DELETE_115"></a>DELETE</h4>
<h5><a id="Headers__Token__token_value_116"></a>Headers - (Token : token_value)</h5>
<h6><a id="Request_118"></a>Request</h6>
<pre><code class="language-sh">{
  <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>
}
</code></pre>
<h6><a id="Response_124"></a>Response</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"success"</span>: <span class="hljs-string">"Succesfully delete the user with email - alice.mcdonnell@apple.com"</span>
}
</code></pre>
<h4><a id="apichange_password_131"></a>/api/change_password</h4>
<h4><a id="POST_132"></a>POST</h4>
<h5><a id="Headers__Token__token_value_133"></a>Headers - (Token : token_value)</h5>
<h6><a id="Request_134"></a>Request</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"email"</span>: <span class="hljs-string">"alice.mcdonnell@apple.com"</span>,
    <span class="hljs-string">"userType"</span>: <span class="hljs-string">"User"</span>,
    <span class="hljs-string">"oldPassword"</span>: <span class="hljs-string">"Hello@123"</span>,
    <span class="hljs-string">"password"</span>: <span class="hljs-string">"Hello@456"</span>
}
</code></pre>
<h6><a id="Response_143"></a>Response</h6>
<pre><code class="language-sh">{
    <span class="hljs-string">"success"</span>: <span class="hljs-string">"Password successfully changed for - alice.mcdonnell@apple.com"</span>
}
</code></pre>
