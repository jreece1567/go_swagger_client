# clean up any old generated code
rm -rf client_core
rm -rf models_core
rm -rf client_secure
rm -rf models_secure
# generate new client-code for the 'core' services
rm swagger.json
ln -s 1_core.json swagger.json
./swagger generate client
mv client client_core
mv models models_core
# generate new client-code for the 'secure' services
rm swagger.json
ln -s 1_secure.json swagger.json
./swagger generate client
mv client client_secure
mv models models_secure
# get rid of temporary symlink
rm swagger.json
# post-generate modifications
echo "Note that client_core/westfield_apis_client.go line 40 needs to be edited to include the correct hostname"
echo "Note that client_secure/westfield_apis_client.go line 40 needs to be edited to include the correct hostname"
echo "Also note that any reference to the label 'WestfieldApis' should be changed to 'WestfieldCoreApis' and 'WestfieldSecureApis', respectively."
#
