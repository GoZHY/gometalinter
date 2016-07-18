package main

func dictAdditions() map[string]string {
	dict := parseWikipediaFormat(additions)
	dict["Dont "] = "Don't "
	dict[" parition"] = " partition"
	dict["artefact "] = "artifact "
	dict["carefull "] = "careful "

	// the following are common in source code comments
	dict[" tood"] = " todo"
	dict["/tood"] = "/todo"
	dict["#tood"] = "#todo"
	dict["*tood"] = "*todo"

	dict = expandCase(dict)

	return dict
}

// arent
var additions = `
entreperure->entrepreneur
entreprenuer->entrepreneur
nuetral->neutral
laready->already
varaible->variable
datbase->database
requrement->requirement
brocoli->brocolli
dependancies->dependencies
emtpy->empty
fandation->foundation
environemnt->environment
verious->various
respository->repository
respositories->repositories
gloabl->global
fragement->fragment
upsteam->upstream
specifing->specifying
overriden->overridden
accesss->access
adderss->address
dashbaord->dashboard
auhtenticate->authenticate
retunred->returned
langauge->language
specifing->specifying
heirachy->hierarchy
authenticor->authenticator
availabale->available
positve->positive
satifies->satisfies
capialized->capitalized
versoin->version
obvioulsy->obviously
fundemental->fundamental
crytopgraphic->cryptographic
appication->application
accending->ascending
consisent->consistent
percision->precision
determinsitic->deterministic
elasped->elapsed
udpated->updated
undescore->underscore
represenation->representation
registery->registry
redundent->redundant
puncutation->punctuation
genrates->generates
finallizes->finalizes
expoch->epoch
equivalant->equivalent
determinsitic->deterministic
normallized->normalized
elasped->elapsed
machiens->machines
demonstates->demonstrates
collumn->column
verical->vertical
refernece->reference
opartor->operator
elimiate->eliminate
coalese->coalesce
extenion->extension
affliated->affiliated
hesistate->hesitate
arrary->array
hunman->human
currate->curate
retuns->returns
interfce->interface
alrorythm->algorithm
credentaisl->credentials
closeing->closing
Constructur->Constructor
Depdending->Depending
Disclamer->Disclaimer
Elimintates->Eliminates
Fowrards->Forwards
Implementor->Implementer
Instalation->Installation
Numerious->Numerous
Specifcation->Specification
Wheter->Whether
aforementioend->aforementioned
annonymouse->anonymous
approstraphe->apostrophe
apporach->approach
aribtrary->arbitrary
asychronous->asynchronous
avaiable->available
cahched->cached
calback->callback
careflly->carefully
commmand->command
compatibilty->compatibility
comptability->compatibility
conatins->contains
conditon->condition
configuraiton->configuration
consitency->consistency
contructed->constructed
contructor->constructor
december->December
declareation->declaration
decomposeion->decomposition
deliviered->delivered
depedencies->dependencies
depedency->dependency
deperecation->deprecation
descriminant->discriminant
diffucult->difficult
documenation->documentation
dyamically->dynamically
embeded->embedded
everwhere->everywhere
exising->existing
explicitely->explicitly
explicity->explicitly
expliots->exploits
exprimental->experimental
extactly->exactly
functionlity->functionality
functtion->function
idiosynchracies->idiosyncrasies
immidiate->immediate
implemention->implementation
implentation->implementation
implicitely->implicitly
implimenation->implementation
incldue->include
incorect->incorrect
incorectly->incorrectly
inferrence->inference
milisecond->millisecond
mimimum->minimum
minimium->minimum
misinterpretting->misinterpreting
momment->moment
muliple->multiple
mulitple->multiple
nubmers->numbers
officiallly->officially
otherhand->other hand
optinally->optimally
ouput->output
outputed->outputted
pacakge->package
packge->package
paramter->parameter
paramters->parameters
paricular->particular
performaces->performances
permisson->permission
precedeed->preceded
precendence->precedence
programattically->programmatically
programmar->programmer
programms->programs
properites->properties
propeties->properties
protototype->prototype
publsih->publish
quuery->query
requried->required
retrived->retrieved
ridiculus->ridiculous
seperator->separator
similarlly->similarly
simplfy->simplify
singals->signals
spanish->Spanish
specifcally->specifically
specifed->specified
specifiy->specify
straitforward->straightforward
subsequant->subsequent
successfuly->successfully
supportied->supported
supression->suppression
synchornously->synchronously
syncronously->synchronously
tutorual->tutorial
unintuive->unintuitive
writting->writing
Euclidian->Euclidean
`