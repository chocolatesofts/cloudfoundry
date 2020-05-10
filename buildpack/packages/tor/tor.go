package tor
import(
	"github.com/chocolatesofts/cloudfoundry/apt"
	"github.com/cloudfoundry/libbuildpack"
	"path/filepath"
)
type Supplier struct{
	AptSupplier			*apt.Supplier
	Logger			    *libbuildpack.Logger
	Config				string
}

func InstallTor(s *Supplier) error {
	s.Logger.Info("Installing Tor.....")
	err:=apt.SingleInstall(s.AptSupplier,"tor","repo")
	if(err!=nil){
		return err
	}
	s.Logger.Info("Tor installed!!!")
	cfile:=configfilename
	if(s.Config!=""){
		cfile=s.Config
	}
	_,err=parseConfig(filepath.Join(s.AptSupplier.Stager.BuildDir(),cfile))
	if(err!=nil){
		return err
	}
	torscript:=`export TOR_PORT_1=58`
	s.AptSupplier.Stager.WriteProfileD("tor.sh",torscript)
	return nil
}