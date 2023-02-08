import {db} from '../../firebase/firebase'
import {doc,getDoc,updateDoc,setDoc} from 'firebase/firestore'


const getUser = async => {
    const userRef = doc(db,"Users",auth.currentUser.uid);
    const userData = (await getDoc(userRef)).date();
    return userData;
}

export {getUser};